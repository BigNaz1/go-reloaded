package reloaded

import (
	"fmt"
)

func Filterstring(s string) string {
	for _, r := range s {
		if r < ' ' || r > '~' {
			fmt.Println("Error: Unrecognized charecter, refer to the user guide for more information")
			return ""
		}
	}
	return s
}
