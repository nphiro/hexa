package crypter

import "time"

type SealOptions struct {
	ID        string
	Subject   string
	Audience  []string
	Duration  *time.Duration
	ExpiresAt *time.Time
}

func NewSealOptions() *SealOptions {
	return new(SealOptions)
}

func (o *SealOptions) WithID(id string) *SealOptions {
	o.ID = id
	return o
}

func (o *SealOptions) WithSubject(subject string) *SealOptions {
	o.Subject = subject
	return o
}

func (o *SealOptions) WithAudience(audience []string) *SealOptions {
	o.Audience = audience
	return o
}

func (o *SealOptions) WithDuration(duration time.Duration) *SealOptions {
	o.Duration = &duration
	return o
}

func (o *SealOptions) WithExpiresAt(expiresAt time.Time) *SealOptions {
	o.ExpiresAt = &expiresAt
	return o
}
