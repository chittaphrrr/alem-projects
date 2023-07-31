package functions

import (
	"strings"
	"unicode"
)

func Convert(content []byte) string {
	text := putSpaces(string(content))
	tempSlice := cleaner(getSlice(text))
	tempSlice = fixArticles(tempSlice)
	tempSlice = quotesExceptions(tempSlice)
	tempSlice = fixMods(tempSlice)
	tempSlice = DoMods(tempSlice)
	tempSlice = quotesAfterMods(tempSlice)
	return fixSpacesAfterMods(tempSlice)
}

func fixSpacesAfterMods(slice []string) string {
	var result string
	for i := range slice {
		result += slice[i]
		switch {
		case i < len(slice)-1 && hasOnlyPunctuation(slice[i+1]):
			continue
		default:
			result += " "
		}
	}
	if result[len(result)-1] == ' ' {
		return result[:len(result)-1]
	}
	return result
}

func putSpaces(s string) string {
	var result string
	runes := []rune(s)
	for i := range runes {
		result += string(runes[i])
		if !isPunctuation(runes[i]) && i < len(runes)-1 && runes[i+1] == '(' {
			result += " "
		} else if isPunctuation(runes[i]) && runes[i] == '(' {
			continue
		} else if isPunctuation(runes[i]) {
			result += " "
		}
	}
	return result
}

func fixMods(slice []string) []string {
	for i := range slice {
		if hasOnlyDigit(slice[i]) && slice[i+1] == ")" {
			slice[i], slice[i+1] = slice[i]+slice[i+1], ""
		}
	}
	return cleaner(slice)
}

func hasOnlyDigit(s string) bool {
	var count int
	for _, r := range s {
		if unicode.IsDigit(r) {
			count++
		}
	}
	if len(s) == count {
		return true
	}
	return false
}

func hasOnlyPunctuation(s string) bool {
	var count int
	for _, r := range s {
		if isPunctuation(r) {
			count++
		}
	}
	if len(s) == count {
		return true
	}
	return false
}

func isPunctuation(r rune) bool {
	if unicode.IsPunct(r) {
		return true
	}
	return false
}

func getString(slice []string) string {
	var result string
	for i := range slice {
		if slice[i] == "\n" {
			result += "\n"
		} else {
			result += slice[i] + " "
		}
	}
	return result
}

func getSlice(s string) []string {
	var temp string
	for _, r := range s {
		if r == '\n' {
			temp += " " + "\n" + " "
		} else if r == ' ' {
			temp += " "
		} else {
			temp += string(r)
		}
	}

	return strings.Split(temp, " ")
}

func cleaner(slice []string) []string {
	result := []string{}
	for _, str := range slice {
		if str == "" || str == " " {
			continue
		}
		result = append(result, str)
	}
	return result
}
