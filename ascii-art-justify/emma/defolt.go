package emma

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func Asciiart() error {
	word := string(os.Args[1])
	text, err := ioutil.ReadFile("data/standard.txt")
	if err != nil {
		return err
	}
	if err = CheckHash("standard"); err != nil {
		return err
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

func Asciiartfs() {
	word := string(os.Args[1])
	text := Fs(os.Args[2])
	if !CheckHash(os.Args[2]) {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
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

func Asciiartjustify() {
	jfy, err := Justify(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	if jfy == "err" {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
		return
	}
	word := string(os.Args[2])
	text := Fs(os.Args[3])
	if !CheckHash(os.Args[3]) {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
		return
	}
	word = strings.ReplaceAll(word, "\n", "\\n")
	wordmas := strings.Split(word, "\\n")
	c := 0
	if jfy == "r" {
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
					if !ThorR(string(text), wordmas[h], i) {
						return
					}
					// Thorjustify(string(text), wordmas[h], i)
				}
			} else if len(wordmas) == 1 && len(g) == 0 {
				return
			} else {
				fmt.Println()
			}
		}
	} else if jfy == "l" {
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
					if !ThorL(string(text), wordmas[h], i) {
						return
					}
					// Thorjustify(string(text), wordmas[h], i)
				}
			} else if len(wordmas) == 1 && len(g) == 0 {
				return
			} else {
				fmt.Println()
			}
		}
	} else if jfy == "c" {
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
					if !ThorC(string(text), wordmas[h], i) {
						return
					}
					// Thorjustify(string(text), wordmas[h], i)
				}
			} else if len(wordmas) == 1 && len(g) == 0 {
				return
			} else {
				fmt.Println()
			}
		}
	} else if jfy == "j" {
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
					if !ThorJ(string(text), wordmas[h], i) {
						return
					}
				}
			} else if len(wordmas) == 1 && len(g) == 0 {
				return
			} else {
				fmt.Println()
			}
		}
	}
	// m, _, _ := GetTerminalSize()
	// fmt.Println(m)
}
