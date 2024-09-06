package httpserver

import (
	"context"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"log/slog"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type Protocol string

const (
	ProtocolHTTP  Protocol = "http"
	ProtocolHTTPS Protocol = "https"
)

func Start(baseCtx context.Context, h http.Handler, opts *ServerOptions) error {
	if opts == nil {
		opts = NewServerOptions()
	}

	// Set the write timeout to 0 if debug is enabled
	h = h2c.NewHandler(h, &http2.Server{})
	h = recoveryMiddleware(h)

	srv := &http.Server{
		Addr:    fmt.Sprintf("localhost:%s", opts.Port),
		Handler: h,
	}

	if !opts.Debug {
		srv.WriteTimeout = opts.WriteTimeout
	}

	protocal := ProtocolHTTP
	if opts.TLSCertBase64 != "" && opts.TLSKeyBase64 != "" {
		protocal = ProtocolHTTPS
		cert, err := decodeTLSCredentials(opts.TLSCertBase64, opts.TLSKeyBase64)
		if err != nil {
			return err
		}
		srv.TLSConfig = &tls.Config{
			Certificates: []tls.Certificate{cert},
			NextProtos:   []string{"h2", "http/1.1"},
		}
	}

	// Create a context that listens for the interrupt signals
	ctx, cancel := signal.NotifyContext(baseCtx, os.Interrupt, os.Kill)
	defer cancel()

	srvErrorChan := make(chan error, 1)
	srvShutdownChan := make(chan struct{})

	// Add graceful shutdown
	go func() {
		select {
		case <-srvErrorChan:
		case <-ctx.Done():
			slog.Info("Server is shutting down...")
			if err := srv.Shutdown(context.Background()); err != nil {
				slog.Error("Server shutdown error", slog.String("error", err.Error()))
			} else {
				slog.Info("Server has been shutdown successfully")
				srvShutdownChan <- struct{}{}
			}
		}
	}()

	slog.Info(fmt.Sprintf("Server is running on %s://%s", protocal, srv.Addr))

	var err error
	if protocal == ProtocolHTTP {
		err = srv.ListenAndServe()
	} else {
		err = srv.ListenAndServeTLS("", "")
	}

	if err != nil && err != http.ErrServerClosed {
		slog.Error("Server encountered an error", slog.String("error", err.Error()))
		srvErrorChan <- err
		return err
	}

	<-srvShutdownChan
	return nil
}

func decodeTLSCredentials(certBase64, keyBase64 string) (tls.Certificate, error) {
	certPem, err := base64.URLEncoding.DecodeString(certBase64)
	if err != nil {
		return tls.Certificate{}, ErrDecodeKeyArgument{Type: "cert"}
	}
	keyPem, err := base64.URLEncoding.DecodeString(keyBase64)
	if err != nil {
		return tls.Certificate{}, ErrDecodeKeyArgument{Type: "key"}
	}
	cert, err := tls.X509KeyPair(certPem, keyPem)
	if err != nil {
		return tls.Certificate{}, ErrInvalidCertKeyPair
	}
	return cert, nil
}

func recoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				slog.Error("Recovered from a panic", slog.Any("error", err))
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
