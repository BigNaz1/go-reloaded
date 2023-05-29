package reloaded

import (
	"strings"
)

func AdvancedPunctuation(input string) string {
	parts := strings.Split(input, "'")

	for i := 1; i < len(parts)-1; i += 2 {
		field := strings.Split(parts[i], "\\s+")
		parts[i] = "'" + strings.TrimSpace(strings.Join(field, " ")) + "'"
	}
	return strings.Join(parts, "")
}
