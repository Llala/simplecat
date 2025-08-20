package util

import "strings"

func ParseTextUtil(sourceText string) []string {
	stringParts := strings.Split(sourceText, ".")
	// Create a new slice to store the trimmed parts
	var cleanedParts []string

	// Iterate over the parts and trim whitespace from each
	for _, part := range stringParts {
		if part != "" {
			cleanedParts = append(cleanedParts, strings.TrimSpace(part))
		}
	}
	return cleanedParts
}
