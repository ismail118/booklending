package token

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

var (
	ErrTokenExpired = errors.New("token expired")
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expires_at"`
}

func NewPayload(email string, duration time.Duration) (*Payload, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        id,
		Email:     email,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return payload, nil
}

func (p *Payload) Validate() error {
	if time.Now().After(p.ExpiredAt) {
		return ErrTokenExpired
	}

	return nil
}
