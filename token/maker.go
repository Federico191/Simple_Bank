package token

import (
	"time"
)

// Maker is an interface to managing token
type Maker interface {
	// CreateToken creates a new token for specific username and password
	CreateToken(username string, duration time.Duration) (string, error)

	// VerifyToken checks if token valid or not
	VerifyToken(token string) (*Payload, error)
}
