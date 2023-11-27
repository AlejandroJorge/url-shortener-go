package logic

import "strings"

func IsValidURL(originalURL string) bool {
	isValid := true
	isValid = isValid && (strings.HasPrefix(originalURL, "http://www.") || strings.HasPrefix(originalURL, "https://www."))
	isValid = isValid && strings.Count(originalURL, ".") >= 2
	return isValid
}
