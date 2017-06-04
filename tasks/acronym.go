// Convert a phrase to its acronym.

package tasks

import (
	"strings"
)

// Acronym from a string
func Acronym(str string) string {
	return strings.Map(func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return r
		default:
			return -1
		}
	}, str)
}
