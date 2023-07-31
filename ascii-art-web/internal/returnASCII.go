package internal

import (
	"strings"
)

func ReturnASCII(inputs []string, banner string) (string, int) {
	symbols, status := getSymbols(banner)
	if status != 200 {
		return "", status
	}

	var res string
	for i := range inputs {
		switch inputs[i] {
		case "\\n":
			res += "\n"
		case "\\t":
			res += "\t"
		default:
			for line := 0; line < 8; line++ {
				for _, r := range inputs[i] {
					res += symbols[r][line]
				}
				if line != 7 {
					res += "\n"
				}
			}
		}
	}
	if !strings.HasSuffix(res, "\n") {
		res += "\n"
	}
	return res, 200
}

func getSymbols(source string) (map[rune][]string, int) {
	content, status := ValidHash(source)
	if status != 200 {
		return map[rune][]string{}, status
	}

	tempstr := strings.ReplaceAll(string(content), "\r", "")

	temp := strings.Split(tempstr[1:], "\n")

	symbols := make(map[rune][]string)

	i := 0
	for key := ' '; key <= '~'; key++ {
		symbols[key] = temp[i : i+8]
		i += 9
	}

	return symbols, 200
}
