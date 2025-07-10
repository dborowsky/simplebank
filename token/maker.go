package token

import (
	"time"
)

type Maker interface {
	// CreateToken creates a new token for the given user with the given duration. It returns the token as a string.
	CreateToken(username string, diration time.Duration) (string, error)
	// VerifyToken verifies the given token and returns the payload if valid. It returns an error otherwise.
	VerifyToken(token string) (*Payload, error)
}
