package internal

import (
	"strings"
)

const source = "\r\n\\n\n !\x22#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~"

func IsValidInput(input []string) (string, int) {
	for _, char := range input[0] {
		if !strings.Contains(source, (string(char))) {
			return "", 400
		}
	}

	switch {
	case len(input) == 2:
		if input[1] == "shadow" {
			return "shadow", 200
		}
		if input[1] == "standard" {
			return "standard", 200
		}
		if input[1] == "thinkertoy" {
			return "thinkertoy", 200
		}
	}

	return "", 400
}
