package functions

import "strings"

func fixArticles(slice []string) []string {
	for i := range slice {
		if i < len(slice)-1 {
			if strings.ToLower(slice[i]) == "a" && StartsWithVowel(slice[i+1]) {
				slice[i] = slice[i] + "n"
			} else if strings.ToLower(slice[i]) == "an" && !StartsWithVowel(slice[i+1]) {
				slice[i] = slice[i][:1]
			}
		}
	}
	return slice
}

func StartsWithVowel(s string) bool {
	if isVowel(s[0]) {
		return true
	}
	return false
}

func isVowel(b byte) bool {
	switch b {
	case 'a', 'o', 'i', 'e', 'u', 'h', 'A', 'O', 'I', 'E', 'U', 'H':
		return true
	}
	return false
}
