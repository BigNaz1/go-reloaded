package reloaded

import (

	"strconv"
	"strings"
	"unicode"
)

func RemoveExtraPunc(input string) string {
	words := strings.Fields(input)
	var builder strings.Builder
	punctuationCount := 0
	maxPunctuation := 3

	for i := 0; i < len(words); i++ {
		word := words[i]

		if strings.HasPrefix(word, "(") && strings.HasSuffix(word, ")") {
			command := strings.ToLower(word[1 : len(word)-1])

			switch command {
			case "cap":
				if i > 0 {
					builder.WriteString(" ")
				}
				i++
				if i < len(words) {
					builder.WriteString(strings.Title(strings.ToLower(words[i])))
				}
			case "up":
				if i > 0 {
					builder.WriteString(" ")
				}
				i++
				if i < len(words) {
					builder.WriteString(strings.ToUpper(words[i]))
				}
			case "low":
				if i > 0 {
					builder.WriteString(" ")
				}
				i++
				if i < len(words) {
					builder.WriteString(strings.ToLower(words[i]))
				}
			default:
				if strings.Contains(command, "cap,") {
					if i > 0 {
						builder.WriteString(" ")
					}
					i++
					if i < len(words) {
						n, err := strconv.Atoi(command[4:])
						if err == nil {
							for j := 0; j < n && i < len(words); j++ {
								if j > 0 {
									builder.WriteString(" ")
								}
								builder.WriteString(strings.Title(strings.ToLower(words[i])))
								i++
							}
							i--
						}
					}
				} else {
					builder.WriteString(" " + word)
				}
			}
		} else {
			if builder.Len() > 0 {
				builder.WriteString(" ")
			}

			for _, r := range word {
				if unicode.IsPunct(r) {
					punctuationCount++
					if punctuationCount <= maxPunctuation {
						builder.WriteRune(r)
					}
				} else {
					punctuationCount = 0
					builder.WriteRune(r)
				}
			}
		}
	}

	return builder.String()
}
