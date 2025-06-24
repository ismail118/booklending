package token

import (
	"fmt"
	"github.com/o1egl/paseto"
	"time"
)

const (
	minKeySize = 32
)

type Paseto struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func NewPaseto(symmetricKey string) (*Paseto, error) {
	if len(symmetricKey) < minKeySize {
		return nil, fmt.Errorf("key must be at least %v charters", minKeySize)
	}

	p := &Paseto{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}

	return p, nil
}

func (p *Paseto) CreateToken(email string, duration time.Duration) (string, error) {
	payload, err := NewPayload(email, duration)
	if err != nil {
		return "", err
	}
	return p.paseto.Encrypt(p.symmetricKey, payload, nil)
}

func (p *Paseto) VerifyToken(token string) (*Payload, error) {
	var payload Payload
	err := p.paseto.Decrypt(token, p.symmetricKey, &payload, nil)
	if err != nil {
		return nil, err
	}

	err = payload.Validate()
	if err != nil {
		return nil, err
	}

	return &payload, nil
}
