package httpserver

import "time"

type ServerOptions struct {
	Port          string
	TLSCertBase64 string
	TLSKeyBase64  string

	Debug        bool
	WriteTimeout time.Duration
}

func NewServerOptions() *ServerOptions {
	return &ServerOptions{
		Port:         "8080",
		WriteTimeout: 10 * time.Second,
	}
}

func (o *ServerOptions) WithPort(port string) *ServerOptions {
	o.Port = port
	return o
}

// WithTLS enables TLS with the provided cert and key which are base64 encoded
func (o *ServerOptions) WithTLS(certBase64, keyBase64 string) *ServerOptions {
	o.TLSCertBase64 = certBase64
	o.TLSKeyBase64 = keyBase64
	return o
}

// WithDebug disables write timeout
func (o *ServerOptions) WithDebug(debug bool) *ServerOptions {
	o.Debug = debug
	return o
}

func (o *ServerOptions) WithWriteTimeout(timeout time.Duration) *ServerOptions {
	o.WriteTimeout = timeout
	return o
}
