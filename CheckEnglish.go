package reloaded

import (
	"fmt"
	//"strings"
	//"unicode"
)

func Filterstring(s string) string {
    for _, r := range s {
        if (r < ' ' || r > '~'){
            fmt.Println("Error: String contains non-English characters")
            return ""
        }
    }
    return s
}