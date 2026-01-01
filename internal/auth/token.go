package auth

import (
	"errors"
)

var (
	ErrInvalidToken = errors.New("invalid authentication token")
	ErrMissingToken = errors.New("missing authentication token")
)

// TokenValidator validates authentication tokens
type TokenValidator struct {
	validTokens map[string]bool
}

// NewTokenValidator creates a new token validator
func NewTokenValidator(tokens []string) *TokenValidator {
	validTokens := make(map[string]bool)
	for _, token := range tokens {
		if token != "" {
			validTokens[token] = true
		}
	}

	return &TokenValidator{
		validTokens: validTokens,
	}
}

// Validate checks if a token is valid
func (tv *TokenValidator) Validate(token string) error {
	if token == "" {
		return ErrMissingToken
	}

	if !tv.validTokens[token] {
		return ErrInvalidToken
	}

	return nil
}

// IsValid returns true if the token is valid
func (tv *TokenValidator) IsValid(token string) bool {
	return tv.validTokens[token]
}
