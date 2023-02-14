// Package helpers hold utility functions and methods
package helpers

import (
	"context"
	"crypto/sha256"
	"errors"
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

// GetContextValue to extract the request-id value from the context passed
func GetContextValue(ctx context.Context, k model.Key) (string, error) {
	if v := ctx.Value(k); v != nil {
		fmt.Println("found value:", v)
		return fmt.Sprintf("%s", v), nil
	}
	return "", errors.New("something-went-wrong")
}
