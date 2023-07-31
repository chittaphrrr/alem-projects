package internal

func Convert(text string, banner string) (string, int) {
	banner, status := IsValidInput([]string{text, banner})
	if status != 200 {
		return "", status
	}

	res, status := ReturnASCII(MySplit(text), banner)

	if status != 200 {
		return "", status
	}

	return res, 200
}
