package reloaded

import (
	"strconv"
	"strings"
    "fmt"
)

func HexatoDec(HextoDec string) string {
    var output int64
    words := strings.Fields(HextoDec)
    for i := 1; i < len(words); i++ {
        if strings.ToLower(words[i]) == "(hex)" {
            number := strings.ToLower(words[i-1])
            words[i] = ""

            var err error
			output, err = strconv.ParseInt(number, 16, 64)
			if err != nil {
				// Print an error message and return an empty string in case of an error
				fmt.Println("There is an error with your input kindly check the user manual before running")
				return ""
			}
            words[i-1] = strconv.FormatInt(output, 10)
        }
    }
    temp := strings.Join(words, " ")
    words = strings.Fields(temp)
    return  strings.Join(words, " ")
}