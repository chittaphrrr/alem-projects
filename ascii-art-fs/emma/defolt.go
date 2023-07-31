package emma

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func Defolt() {
	word := string(os.Args[1])
	text, err := ioutil.ReadFile("data/standard.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	if !CheckHash("standard") {
		fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard\nYour BANNER  has been changed")
		return
	}
	word = strings.ReplaceAll(word, "\n", "\\n")
	wordmas := strings.Split(word, "\\n")

	c := 0
	for _, i := range wordmas {
		if len(i) > 0 {
			c++
		}
	}
	if c == 0 {
		for i := 0; i < len(wordmas)-1; i++ {
			fmt.Println()
		}
		return
	}

	for h, g := range wordmas {
		if len(g) > 0 {
			for i := 1; i <= 8; i++ {
				Thor(string(text), wordmas[h], i)
			}
		} else if len(wordmas) == 1 && len(g) == 0 {
			return
		} else {
			fmt.Println()
		}
	}
}
