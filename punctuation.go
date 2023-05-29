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
			if i > 0 && unicode.IsSpace(rune(input[i-1])) && (char != '\'' || (char == '\'' && i+1 < len(input) && rune(input[i+1]) != ' ')) {
				newStr := builder.String()
				newStr = newStr[:len(newStr)-1]
				builder.Reset()
				builder.WriteString(newStr)
			}

			builder.WriteRune(char)
			i++

			if (char == '\'' || char == ',') && i < len(input) && !unicode.IsSpace(rune(input[i])) && !unicode.IsPunct(rune(input[i])) {
				builder.WriteRune(' ')
			}
		} else {
			builder.WriteRune(char)
			i++
		}
	}

	return builder.String()
}
