package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// Different types of error returned by the VerifyToken function
var (
	ErrInvalidToken = fmt.Errorf("token is invalid")
	// ErrExpiredToken = fmt.Errorf("token has expired")
	ErrExpiredToken = fmt.Errorf("token has invalid claims: token is expired")

)

type Payload struct {
	ID               uuid.UUID `json:"id"`
	Username         string    `json:"username"`
	RegisteredClaims jwt.RegisteredClaims
}



func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:       tokenID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			Issuer: "Mriganka Dhar",
		},
	}
	return payload, nil
}

// // Valid checks if the token payload is valid or not
// func Valid(payload *Payload) error {
// 	if time.Now().After(payload.ExpiredAt) {
// 		return ErrExpiredToken
// 	}
// 	return nil
// }

// GetAudience implements jwt.Claims.
func (payload *Payload) GetAudience() (jwt.ClaimStrings, error) {
	return payload.RegisteredClaims.Audience,nil;
}

// GetExpirationTime implements jwt.Claims.
func (payload *Payload) GetExpirationTime() (*jwt.NumericDate, error) {
	return payload.RegisteredClaims.ExpiresAt,nil;
}

// GetIssuedAt implements jwt.Claims.
func (payload *Payload) GetIssuedAt() (*jwt.NumericDate, error) {
	return payload.RegisteredClaims.IssuedAt,nil;
}

// GetIssuer implements jwt.Claims.
func (payload *Payload) GetIssuer() (string, error) {
	return payload.RegisteredClaims.Issuer,nil;
}

// GetNotBefore implements jwt.Claims.
func (payload *Payload) GetNotBefore() (*jwt.NumericDate, error) {
	return payload.RegisteredClaims.NotBefore,nil;

}

// GetSubject implements jwt.Claims.
func (payload *Payload) GetSubject() (string, error) {
	return payload.RegisteredClaims.Subject,nil;
}

// Valid implements jwt.Claims.
func (payload *Payload) Valid() error {
	if time.Now().After(payload.RegisteredClaims.ExpiresAt.Time) {
		return ErrExpiredToken
	}
	return nil
}
