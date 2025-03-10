package reloaded

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func TransformText(input string) string {
	words := strings.Fields(input)

	noSpace := regexp.MustCompile(`\((low|up|cap)\d+\)`)
	noComma := regexp.MustCompile(`\((low|up|cap),\d+\)`)
	noBoth := regexp.MustCompile(`\((low|up|cap)`)

	for i := 0; i < len(words); i++ {
		if noSpace.MatchString(words[i]) || noComma.MatchString(words[i]) ||
			noBoth.MatchString(words[i]) && !containsSubstring(words[i], ",") {
			fmt.Println("you have an Invalid marker. A marker should contain both a comma and a space.")
			os.Exit(0)
		}

		if strings.HasPrefix(words[i], "(low,") || strings.HasPrefix(words[i], "(up,") || strings.HasPrefix(words[i], "(cap,") {

			parts := strings.Split(words[i], ",")
			if len(parts) == 2 {
				number, err := strconv.Atoi(words[i+1][:len(words[i+1])-1])
				if number > i {
					fmt.Println("your number inside the up/low/cap marker is larger than number of words before it")
					os.Exit(0)
				} else if err == nil && i-number >= 0 {

					if strings.HasPrefix(words[i], "(low,") {
						for j := i - number; j < i; j++ {
							words[j] = strings.ToLower(words[j])
						}
					} else if strings.HasPrefix(words[i], "(up,") {
						for j := i - number; j < i; j++ {
							words[j] = strings.ToUpper(words[j])
						}
					} else if strings.HasPrefix(words[i], "(cap,") {
						for j := i - number; j < i; j++ {
							words[j] = strings.Title(words[j])
						}
					}

					words[i+1] = ""
					words[i] = ""
					words = append(words[:i], words[i+1:]...)
					i -= number
					temp := strings.Join(words, " ")
					words = strings.Fields(temp)
				}
			}
		}
	}

	return strings.Join(words, " ")
}
