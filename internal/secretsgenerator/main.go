package secretsgenerator

import (
	"crypto/sha256"
	"encoding/base64"

	"github.com/google/uuid"
)

func GenerateState() []byte {
	r1 := uuid.New()
	r2 := uuid.New()

	return []byte(r1.String() + r2.String())
}

func GenerateVerifier(state []byte) string {
	return base64URLEncode(string(state))
}

func GenerateChallenge(verifier string) string {
	hash := generateSHA256Hash([]byte(verifier))

	return base64URLEncode(string(hash))
}

func generateSHA256Hash(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

func base64URLEncode(data string) string {
	return base64.RawURLEncoding.EncodeToString([]byte(data))
}
