package reloaded

import "strings"

func ReplaceSymbol(input string) string {
    return strings.Replace(input, "‘", "'", -1)
}