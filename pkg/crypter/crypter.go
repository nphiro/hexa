package crypter

import (
	"crypto"
	"encoding/base64"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CrypterClient struct {
	options *ClientOptions

	signingMethod jwt.SigningMethod
	signKey       crypto.PrivateKey
	verifyKey     crypto.PublicKey
}

type claimsData[T any] struct {
	Data T `json:"data"`
	jwt.RegisteredClaims
}

func New(privateKeyBase64, publicKeyBase64 string, opts *ClientOptions) (c *CrypterClient, err error) {
	if opts == nil {
		opts = NewClientOptions()
	}

	privateKey, err := base64.URLEncoding.DecodeString(privateKeyBase64)
	if err != nil {
		err = ErrDecodeKeyArgument{Type: "private"}
		return
	}
	publicKey, err := base64.URLEncoding.DecodeString(publicKeyBase64)
	if err != nil {
		err = ErrDecodeKeyArgument{Type: "public"}
		return
	}

	c = &CrypterClient{options: opts, signingMethod: jwt.SigningMethodEdDSA}
	if c.signKey, err = jwt.ParseEdPrivateKeyFromPEM(privateKey); err != nil {
		c = nil
		err = ErrParseKeyArgument{Type: "private"}
		return
	}
	if c.verifyKey, err = jwt.ParseEdPublicKeyFromPEM(publicKey); err != nil {
		c = nil
		err = ErrParseKeyArgument{Type: "public"}
		return
	}
	return
}

func Seal[T any](c *CrypterClient, data T, opts *SealOptions) (string, error) {
	claims := &claimsData[T]{
		data,
		getRegisterClaims(c.options, opts),
	}
	return jwt.NewWithClaims(c.signingMethod, claims).SignedString(c.signKey)
}

func Open[T any](c *CrypterClient, token string) (T, jwt.RegisteredClaims, error) {
	claims := &claimsData[T]{}
	_, err := jwt.ParseWithClaims(token, claims, func(_ *jwt.Token) (any, error) {
		return c.verifyKey, nil
	})
	if err != nil {
		return claims.Data, jwt.RegisteredClaims{}, ErrInvalidToken
	}
	return claims.Data, claims.RegisteredClaims, nil
}

func getRegisterClaims(base *ClientOptions, extended *SealOptions) jwt.RegisteredClaims {
	claims := jwt.RegisteredClaims{
		Issuer:    base.Issuer,
		Subject:   base.Subject,
		Audience:  base.Audience,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
	}

	if base.Duration > 0 {
		claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(base.Duration))
	}

	if extended == nil {
		return claims
	}

	if extended.ID != "" {
		claims.ID = extended.ID
	}
	if extended.Subject != "" {
		claims.Subject = extended.Subject
	}
	if extended.Audience != nil {
		claims.Audience = extended.Audience
	}
	if extended.Duration != nil {
		if *extended.Duration > 0 {
			claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(*extended.Duration))
		} else {
			claims.ExpiresAt = nil
		}
	} else if extended.ExpiresAt != nil {
		claims.ExpiresAt = jwt.NewNumericDate(*extended.ExpiresAt)
	}

	return claims
}
