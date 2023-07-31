package functions

import (
	"errors"
	"strings"
)

func ValidateArguments(args []string) (bool, error) {
	switch {
	case len(args) < 3:
		return false, errors.New("not enough arguments")
	case len(args) > 3:
		return false, errors.New("too many arguments")
	case len(args[1]) < 5 || len(args[2]) < 5:
		return false, errors.New("something's wrong with the files")
	case strings.HasPrefix(args[2], "functions/") || args[2] == "main.go" || args[2] == "go.mod" || args[2] == "README.md":
		return false, errors.New("you are trying to re-write a source file!")
	case !strings.HasSuffix(args[1], ".txt") || !strings.HasSuffix(args[2], ".txt"):
		return false, errors.New("the files are not txt files")
	case args[1] == args[2]:
		return false, errors.New("input file is the same as the output file")
	default:
		return true, nil
	}
}
