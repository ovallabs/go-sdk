package helpers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetSignatureFromReferenceAndPubKey(t *testing.T) {
	tests := []struct {
		reference string
		publicKey string
		expected  string
	}{
		{
			reference: "ref123",
			publicKey: "IYtqe0xG0voYzbPhUEtTIEKyKj4Keq0O",
			expected:  "8c7ac3be5b0a6bee420d573325904f634f6a50ab6a5a4544b0ff8a1c7327264e",
		},
		{
			reference: "ref0091",
			publicKey: "ANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCg",
			expected:  "8cfc061f6b90ccab9ed8f36fa8791e0eff4d3111f8858230bf063970f547ade0",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("case_%d", i), func(t *testing.T) {
			got := GetSignatureFromReferenceAndPubKey(test.reference, test.publicKey)
			require.Equal(t, test.expected, got)
		})
	}
}
