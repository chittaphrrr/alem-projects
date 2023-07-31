package internal

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func ValidHash(banner string) (string, int) {

	standardHash := "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf"
	thinkertoyHash := "64285e4960d199f4819323c4dc6319ba34f1f0dd9da14d07111345f5d76c3fa3"
	shadowHash := "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73"

	currentBanner, err := os.ReadFile("internal/maps/" + banner + ".txt")
	if err != nil {
		return "", 500
	}
	currentHash := fmt.Sprintf("%x", sha256.Sum256(currentBanner))

	if banner == "standard" {
		if standardHash != currentHash {
			return "", 500
		}
	}

	if banner == "thinkertoy" {
		if thinkertoyHash != currentHash {
			return "", 500
		}
	}

	if banner == "shadow" {
		if shadowHash != currentHash {
			return "", 500
		}
	}

	return string(currentBanner), 200
}
