package logic

import (
	"crypto/sha256"
	"encoding/base64"
	"strings"
)

func TransformURL(originalURL string) string {
	hash := sha256.Sum256([]byte(originalURL))

	base64Hash := base64.URLEncoding.EncodeToString(hash[:])

	alphanumericHash := filterAlphaNumeric(base64Hash)

	shortened := alphanumericHash[:12]

	return shortened
}

func filterAlphaNumeric(input string) string {
	var result strings.Builder
	for _, char := range input {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			result.WriteRune(char)
		}
	}
	return result.String()
}
