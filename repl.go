package main

import (
	"strings"
)

func cleanInput(text string) []string {
	cleanText := strings.TrimSpace(strings.ToLower(text))

	return strings.Fields(cleanText)
}
