package reloaded

import (
	"strings"
)

func Advancedpunctuation(input string) string {
	parts := strings.Split(input, "'")

	for i := 1; i < len(parts)-1; i += 2 {
		field := strings.Split(parts[i], "\\s+")

		// Join words with a space and trim spaces
		parts[i] = "'" + strings.TrimSpace(strings.Join(field, " ")) + "'"
	}
	return strings.Join(parts, "")
}

//func Advanedpunctuation(input string) string {
//	input = strings.Replace(input, " ' ", " ", -1)
//	words := strings.Split(input, " ")
//
//	var punctuatedString string
//	if words[0] == "'" {
//		punctuatedString = "'"
//	}
//
//	for i := 0; i < len(words); i++ {
//		if i > 0 && words[i] == "'" {
//			punctuatedString += " " + words[i-1] + "'"
//			i++
//		} else {
//			punctuatedString += words[i] + " "
//		}
//	}
//	return strings.TrimSpace(punctuatedString)
//}
