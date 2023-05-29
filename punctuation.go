package reloaded

import (
	"strings"
	"unicode"
)

func AdjustPunctuationSpacings(input string) string {
	var builder strings.Builder
	i := 0

	for i < len(input) {
		char := rune(input[i])

		if unicode.IsPunct(char) {
			if i > 0 && unicode.IsSpace(rune(input[i-1])) {
				newStr := builder.String()
				newStr = newStr[:len(newStr)-1]
				builder.Reset()
				builder.WriteString(newStr)
			}

			for i < len(input) && unicode.IsPunct(rune(input[i])) {
				builder.WriteRune(rune(input[i]))
				i++
			}

			if i < len(input) && !unicode.IsSpace(rune(input[i])) {
				builder.WriteRune(' ')
			}
		} else {
			builder.WriteRune(char)
			i++
		}
	}

	return builder.String()
}
