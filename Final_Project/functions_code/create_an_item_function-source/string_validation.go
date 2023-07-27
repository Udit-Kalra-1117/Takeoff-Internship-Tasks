package helloworld

import (
	"regexp"
	"strings"
	"unicode"
)

// function to check if string parameters are passed as value
func IsString(v string) bool {
	for _, ch := range v {
		if !unicode.IsLetter(ch) && !unicode.IsSpace(ch) {
			return false
		}
	}
	return true
}

// normalizeName normalizes the name parameter by removing leading/trailing spaces and converting consecutive spaces to a single space
func NormalizeName(s string) string {
	// Remove leading/trailing spaces
	trimmed := strings.TrimSpace(s)
	// Replace consecutive spaces with a single space
	normalized := regexp.MustCompile(`\s+`).ReplaceAllString(trimmed, " ")
	return normalized
}
