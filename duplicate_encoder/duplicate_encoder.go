package kata

import "strings"

func DuplicateEncode(word string) string {
	newString := ""
	for _, ch := range word {
		if strings.Count(strings.ToLower(word), strings.ToLower(string(ch))) > 1 {
			newString += ")"
		} else {
			newString += "("
		}
	}

	return newString
}
