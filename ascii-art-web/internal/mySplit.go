package internal

import (
	"regexp"
)

func MySplit(input string) []string {
	var new_slice []string
	re := regexp.MustCompile(`\n|\\n|\\t|[ -~]`)
	slice := re.FindAllString(input, -1)
	var word string
	for _, v := range slice {
		switch {
		case v != "\\n" && v != "\n" && v != "\\t":
			word += v
		case v == "\\t" || v == "\t":
			word += "    "
		case v == "\n":
			if word != "" {
				new_slice = append(new_slice, word)
			}
			new_slice = append(new_slice, "\\n")
			word = ""
		case word != "":
			new_slice = append(new_slice, word)
			new_slice = append(new_slice, v)
			word = ""
		default:
			new_slice = append(new_slice, v)
		}
	}
	if word != "" {
		new_slice = append(new_slice, word)
	}
	return new_slice
}
