package emma

import (
	"errors"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func GetTerminalSize() (int, int, error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	output, err := cmd.Output()
	if err != nil {
		return 0, 0, err
	}

	sizeStr := strings.TrimSpace(string(output))
	sizeParts := strings.Split(sizeStr, " ")
	if len(sizeParts) != 2 {
		return 0, 0, errors.New("unexpected output format")
	}

	width, err := strconv.Atoi(sizeParts[1])
	if err != nil {
		return 0, 0, err
	}

	height, err := strconv.Atoi(sizeParts[0])
	if err != nil {
		return 0, 0, err
	}

	return width, height, nil
}

func Justify(arg string) (string, error) {
	if arg == "--align=right" {
		return "r", nil
	} else if arg == "--align=left" {
		return "l", nil
	} else if arg == "--align=center" {
		return "c", nil
	} else if arg == "--align=justify" {
		return "j", nil
	}

	return "", errors.New("wrong flag for alignment")
}
