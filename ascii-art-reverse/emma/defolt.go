package emma

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func Defolt() {
	word := string(os.Args[1])
	if strings.Contains(word, "-reverse") {
		Chain()
		return
	}
	if strings.Contains(word, "--output") {
		log.Fatal("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		return
	}
	text, err := ioutil.ReadFile("data/standard.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	if !CheckHash("standard") {
		log.Fatal("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
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
