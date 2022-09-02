package helpers

import (
	"crypto/sha256"
	"fmt"
)

func GetSignatureFromReferenceAndPubKey(reference, publicKey string) string {
	concat := fmt.Sprintf("%s%s", publicKey, reference)
	hash := sha256.New()
	hash.Write([]byte(concat))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
