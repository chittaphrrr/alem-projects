package functions

import "strings"

func quotesAfterMods(slice []string) []string {
	words := cleaner(slice)
	words = fixQuotes(words, '"')
	words = fixQuotes(words, '\'')
	words = fixQuotes(words, '‘')
	words = fixQuotes(words, '`')
	words = cleaner(words)
	words = trimSpace(words)
	return words
}

func quotesExceptions(slice []string) []string {
	for i := range slice {
		if i < len(slice)-1 && i > 0 && strings.HasSuffix(slice[i-1], "n") && slice[i] == "'" && strings.HasPrefix(slice[i+1], "t") {
			slice[i-1], slice[i], slice[i+1] = "", slice[i-1]+slice[i]+slice[i+1], ""
		} else if i < len(slice)-1 && (strings.HasSuffix(slice[i], "n'") && strings.HasPrefix(slice[i+1], "t")) {
			slice[i], slice[i+1] = slice[i]+slice[i+1], ""
		} else if i < len(slice)-1 && strings.HasSuffix(slice[i], "n") && strings.HasPrefix(slice[i+1], "'t") {
			slice[i], slice[i+1] = slice[i]+slice[i+1], ""
		}
	}
	return cleaner(slice)
}

func fixQuotes(slice []string, r rune) []string {
	var firstQuotes []int
	var secondQuotes []int
	var temp bool
	for i := range slice {
		if slice[i] == string(r) && !temp {
			firstQuotes = append(firstQuotes, i)
			temp = true
		} else if slice[i] == string(r) && temp {
			secondQuotes = append(secondQuotes, i)
			temp = false
		}
	}
	for _, i := range firstQuotes {
		if i < len(slice)-1 {
			if slice[i+1] != "\n" {
				slice[i+1] = slice[i] + slice[i+1]
				slice[i] = ""
			} else {
				if i < len(slice)-1 {
					i++
				}
			}
		}
	}
	for _, i := range secondQuotes {
		if i > 0 {
			slice[i-1] += slice[i]
			slice[i] = ""
		}
	}
	return slice
}

func trimSpace(slice []string) []string {
	var result string
	var runes []rune = []rune(getString(slice))
	for i := range runes {
		switch {
		case i > 0 && runes[i] == ' ' && runes[i-1] == '\n':
			continue
		default:
			result += string(runes[i])
			continue
		}
	}
	return strings.Split(result, " ")
}
