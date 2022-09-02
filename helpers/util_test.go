package helpers

import (
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
			expected:  "7097832561ba902beab24076564d53beea425a539046733578ed6e7e2510eb30",
		},
	}

	for _, test := range tests {
		t.Run("testName", func(t *testing.T) {
			got := GetSignatureFromReferenceAndPubKey(test.reference, test.publicKey)
			require.Equal(t, got, test.expected)
		})
	}
}
