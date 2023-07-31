package emma

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func reverseAscii(str string) string {
	return ""
}

func Chain() {
	reverseFile := flag.String("reverse", "", "your file name")
	flag.Usage = func() {
		fmt.Printf("Usage: go run . [OPTION]\n\nEX: go run . --reverse=<fileName>")
		return
	}
	flag.Parse()

	if *reverseFile == "" {
		flag.Usage()
		return
	}

	fileExt := filepath.Ext(*reverseFile)
	if fileExt != ".txt" {
		log.Fatal("Invalid file extension. File name must end with '.txt'.\nUsage: go run . [OPTION]\n\nEX: go run . --reverse=<fileName>")
	}

	fileName := strings.Trim(*reverseFile, "<>")
	old, _ := os.ReadFile(fileName)
	updateFile(fileName)

	if strings.Contains(fileName, "data") {
		log.Fatal("Invalid file name. File name cannot contain the word 'data'.")
	}
	template, e := ReadLine("data/standard.txt")
	if e {
		return
	}
	masTemplate := onesimbol(template)
	fileW, e := ReadLine(fileName)
	if e {
		return
	}

	str := ""
	for len(fileW[0]) > 0 {
		for i, v := range masTemplate {
			if search(v, fileW) {
				str = str + string(rune(i+32))
				fileW = removeLetter(len(v[0]), fileW)
			}
		}
	}
	fmt.Println(str)
	err := ioutil.WriteFile(fileName, old, 0644)
	if err != nil {
		fmt.Println("Ошибка при перезаписи файла:", err)
		return
	}
}

func removeLetter(length int, word []string) []string {
	for i, v := range word[0 : len(word)-1] {
		// fmt.Println(v[length:])
		word[i] = v[length:]
	}
	return word
}

func search(letter, word []string) bool {
	found := true
	if len(letter[0]) > len(word[0]) {
		return false
	}
	for i, v := range word[0 : len(word)-1] {
		if letter[i] != v[:len(letter[i])] {
			found = false
		}
	}
	return found
}

func onesimbol(template []string) [][]string {
	var result [][]string
	for i := 0; i < len(template)-1; i = i + 9 {
		temp := template[i+1 : i+9]
		result = append(result, temp)
	}
	// fmt.Println(len(result))
	return result
}

func ReadLine(o string) ([]string, bool) {
	file := fmt.Sprintf(o)
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return nil, true
	}
	lines := strings.Split(string(bytes), "\n")

	return lines, false
}

func updateFile(filename string) error {
	// Читаем содержимое файла
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	segnew, err := ioutil.ReadFile("data/newline.txt")
	if err != nil {
		return err
	}
	newlines := strings.Split(string(segnew), "\n")
	// Разбиваем содержимое на строки
	lines := strings.Split(string(content), "\n")
	pines := []string{}
	numdelstart := removeLeadingEmpty(lines)
	// fl:=true
	for _, i := range lines {
		if len(i) > 0{
			pines = append(pines, i)
		}
	}
	// for _, i := range lines {
	// 	if len(i)!=0{
	// 		fl=false
	// 	}
	// 	if len(i) > 0 || !fl{
	// 		pines = append(pines, i)
	// 	}
	// }

	lines = pines
	if numdelstart > 0 {
		for i := 0; i < numdelstart; i++ {
			for j := 0; j < 8; j++ {
				lines[j] = newlines[j] + newlines[8+j] + lines[j]
			}
		}
	}
	// Проверяем количество строк
	if len(lines) > 8 {
		for h, i := range lines {
			if h%8 == 0 && h != 0 {
				for j := 0; j < 8; j++ {
					lines[j] += newlines[j] + newlines[8+j]
				}
			}
			if h >= 8 {
				c := h % 8
				lines[c] += i
				lines[h] = ""
			}
		}

		mergedLines := strings.Join(lines[:8], "\n")
		err := ioutil.WriteFile(filename, []byte(mergedLines), 0644)
		if err != nil {
			return err
		}
		// fmt.Println("Файл успешно обновлен.")
	} else {
		// fmt.Println("Количество строк в файле меньше или равно 8. Файл не нуждается в обновлении.")
	}

	return nil
}

func removeLeadingEmpty(arr []string) int {
	leadingEmptyCount := 0
	for _, str := range arr {
		if str == "" {
			leadingEmptyCount++
		} else {
			break
		}
	}
	return leadingEmptyCount
}
