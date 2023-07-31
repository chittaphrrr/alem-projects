package main

import (
	"ascii-art/emma"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=fileName.txt something standard")
		}
	}()
	if len(os.Args) == 3 && strings.Contains(os.Args[1], "-output=") {
		emma.Output2()
		return
	}

	if len(os.Args) == 2 {
		emma.Defolt()
		return
	}
	if len(os.Args) == 4 {
		emma.Output()
		return
	}
	if len(os.Args) > 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=fileName.txt something standard")
		return
	}

	word := string(os.Args[1])
	text := emma.Fs(os.Args[2])
	if !emma.CheckHash(os.Args[2]) {
		fmt.Println("problem with hash\nUsage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=fileName.txt something standard")
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
				emma.Thor(string(text), wordmas[h], i)
			}
		} else if len(wordmas) == 1 && len(g) == 0 {
			return
		} else {
			fmt.Println()
		}
	}
}

// [hello]\n[]\n[]\n[world]
