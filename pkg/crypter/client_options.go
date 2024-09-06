package crypter

import (
	"time"
)

type ClientOptions struct {
	Issuer   string
	Subject  string
	Audience []string
	Duration time.Duration
}

func NewClientOptions() *ClientOptions {
	return &ClientOptions{
		Issuer: "crypter",
	}
}

func (o *ClientOptions) WithIssuer(issuer string) *ClientOptions {
	o.Issuer = issuer
	return o
}

func (o *ClientOptions) WithSubject(subject string) *ClientOptions {
	o.Subject = subject
	return o
}

func (o *ClientOptions) WithAudience(audience []string) *ClientOptions {
	o.Audience = audience
	return o
}

func (o *ClientOptions) WithDuration(duration time.Duration) *ClientOptions {
	o.Duration = duration
	return o
}
