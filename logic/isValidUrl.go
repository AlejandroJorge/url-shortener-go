package logic

import "strings"

func IsValidURL(originalURL string) bool {
	return strings.HasPrefix(originalURL, "http://") || strings.HasPrefix(originalURL, "https://")
}
