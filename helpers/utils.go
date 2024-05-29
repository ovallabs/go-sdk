// Package helpers hold utility functions and methods
package helpers

import (
	"context"
	"crypto/sha256"
	"fmt"

	"github.com/ovalfi/go-sdk/model"
)

// GetSignatureFromReferenceAndPubKey returns the string equivalent of a SHA256 hash on reference and public key
func GetSignatureFromReferenceAndPubKey(reference, publicKey string) string {
	concat := fmt.Sprintf("%s%s", publicKey, reference)
	hash := sha256.New()
	hash.Write([]byte(concat))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

// GetRequestID to extract request-id from context
func GetRequestID(ctx context.Context) string {
	if rID := ctx.Value(model.RequestIDContextKey); rID != nil {
		return rID.(string)
	}

	return ""
}

// GetPointerString get string pointer
func GetPointerString(s string) *string {
	return &s
}

// GetPointerInt get int pointer
func GetPointerInt(i int) *int {
	return &i
}

// GetPointerFloat64 get float64 pointer
func GetPointerFloat64(f float64) *float64 {
	return &f
}
