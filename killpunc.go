package reloaded

import (
	"strings"
)

func RemoveSpacesAroundQuotess(s string) string {
	firstApostrophe := strings.Index(s, "'")
	secondApostrophe := strings.LastIndex(s, "'")

	if firstApostrophe == -1 || secondApostrophe == -1 {
		return s
	}

	beforeFirstApostrophe := strings.TrimRight(s[:firstApostrophe], " ")
	space := " "
	if s[firstApostrophe-1] != '.' && s[firstApostrophe-1] != ',' {
		beforeFirstApostrophe += space  
	}
	insideQuotes := s[firstApostrophe:secondApostrophe+1]
	afterSecondApostrophe := strings.TrimLeft(s[secondApostrophe+1:], " ")
	if s[secondApostrophe] == '.' || s[secondApostrophe] == ','  {
		return  beforeFirstApostrophe + insideQuotes + afterSecondApostrophe[:1]
	}
	return  beforeFirstApostrophe + insideQuotes + afterSecondApostrophe  
}