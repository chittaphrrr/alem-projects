package emma

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Output() {
	if !CheckHash(os.Args[3]) {
		fmt.Println("problem with hash\nUsage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=fileName.txt something standard")
		return
	}
	text := Pechat()
	outputFile := flag.String("output", "", "your file name")
	flag.Usage = func() {
		fmt.Printf("Usage: go run . [OPTION] [STRING] [BANNER]\n\n")
		fmt.Printf("EX: go run . --output=fileName.txt something standard\n")
		// flag.PrintDefaults()
		return
	}
	flag.Parse()

	if *outputFile == "" {
		log.Fatal("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=fileName.txt something standard")
	}
	fileExt := filepath.Ext(*outputFile)
	if fileExt != ".txt" {
		log.Fatal("Invalid file extension. File name must end with '.txt'.\n\ngo run . --output=fileName.txt something standard")
	}

	fileName := strings.Trim(*outputFile, "<>")

	if strings.Contains(fileName, "data") {
		log.Fatal("Недопустимое имя файла. Имя файла не может содержать слово 'data'.")
		log.Fatal("Invalid file extension. File name must end with '.txt'.\n\ngo run . --output=fileName.txt something standard")
	}

	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=fileName.txt something standard", err)
	}
	_, err = file.WriteString(text)
	if err != nil {
		log.Fatal("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=fileName.txt something standard", err)
	}
	defer file.Close()
}

func Output2() {
	text := Pechat2()
	outputFile := flag.String("output", "", "your file name")
	flag.Usage = func() {
		fmt.Printf("Usage: go run . [OPTION] [STRING] [BANNER]\n\n")
		fmt.Printf("EX: go run . --output=fileName.txt something standard\n")
		// flag.PrintDefaults()
		return
	}
	flag.Parse()

	if *outputFile == "" {
		log.Fatal("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=fileName.txt something standard")
	}
	fileExt := filepath.Ext(*outputFile)
	if fileExt != ".txt" {
		log.Fatal("Invalid file extension. File name must end with '.txt'.\n\ngo run . --output=fileName.txt something standard")
	}

	fileName := strings.Trim(*outputFile, "<>")

	if strings.Contains(fileName, "data") {
		log.Fatal("Недопустимое имя файла. Имя файла не может содержать слово 'data'.")
		log.Fatal("Invalid file extension. File name must end with '.txt'.\n\ngo run . --output=fileName.txt something standard")
	}

	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=fileName.txt something standard", err)
	}
	_, err = file.WriteString(text)
	if err != nil {
		log.Fatal("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=fileName.txt something standard", err)
	}
	defer file.Close()
}

func Pechat() string {
	sozder := ""
	word := string(os.Args[2])
	text := Fs(os.Args[3])
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
			sozder += "\n"
		}
		return ""
	}

	for h, g := range wordmas {
		if len(g) > 0 {
			for i := 1; i <= 8; i++ {
				sozder += ThorStr(string(text), wordmas[h], i) + "\n"
			}
		} else if len(wordmas) == 1 && len(g) == 0 {
			return ""
		} else {
			sozder += "\n"
		}
	}
	// file, err := os.Create(fileName)
	// if err != nil {
	// 	log.Fatal("Не удалось создать файл: ", err)
	// }

	// _, err = file.WriteString(sozder)
	// if err != nil {
	// 	log.Fatal("Не удалось записать текст в файл: ", err)
	// }
	// defer file.Close()
	return sozder
}

func Pechat2() string {
	sozder := ""
	word := string(os.Args[2])
	text := Fs("standard")
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
			sozder += "\n"
		}
		return ""
	}

	for h, g := range wordmas {
		if len(g) > 0 {
			for i := 1; i <= 8; i++ {
				sozder += ThorStr(string(text), wordmas[h], i) + "\n"
			}
		} else if len(wordmas) == 1 && len(g) == 0 {
			return ""
		} else {
			sozder += "\n"
		}
	}
	return sozder
}

func ThorStr(ascistr, args string, qatar int) string {
	ascistr = strings.ReplaceAll(ascistr, "\r", "")
	ascimas := strings.Split(ascistr, "\n")
	str := ""
	for _, elem := range []rune(args) {
		str += ascimas[(int(elem)-32)*9+qatar]
	}
	return str
}
