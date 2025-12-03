package signer

import "errors"

// Sentinel errors for common error conditions.
var (
	// ErrInvalidPrivateKey is returned when a private key cannot be decoded or parsed.
	ErrInvalidPrivateKey = errors.New("invalid private key")
)
