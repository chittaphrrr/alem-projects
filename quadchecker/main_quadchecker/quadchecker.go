package main

import (
	"fmt"
	"os"
)

func intToString(n int) string {
	var answer string
	for n > 0 {
		answer += string(n%10 + '0')
		n /= 10
	}
	var reverseAnswer string
	for i := len(answer) - 1; i >= 0; i-- {
		reverseAnswer += string(answer[i])
	}
	return reverseAnswer
}

func Quad(x, y int, ch map[rune]byte) (str string) {
	if 0 < x && 0 < y {
		for row := 1; row <= y; row++ {
			for col := 1; col <= x; col++ {
				if row == 1 {
					if col == 1 {
						str += string(ch['⇖'])
					} else if col == x {
						str += string(ch['⇗'])
					} else {
						str += string(ch['⇑'])
					}
				} else if row == y {
					if col == 1 {
						str += string(ch['⇙'])
					} else if col == x {
						str += string(ch['⇘'])
					} else {
						str += string(ch['⇓'])
					}
				} else {
					if col == 1 {
						str += string(ch['⇐'])
					} else if col == x {
						str += string(ch['⇒'])
					} else {
						str += string(ch['🅂'])
					}
				}
			}
			str += string('\n')
		}
	}
	return
}

func QuadA(x, y int) string {
	ch := map[rune]byte{
		'⇖': 'o',
		'⇗': 'o',
		'⇘': 'o',
		'⇙': 'o',
		'⇐': '|',
		'⇒': '|',
		'⇑': '-',
		'⇓': '-',
		'🅂': ' ',
	}
	return Quad(x, y, ch)
}

func QuadB(x, y int) string {
	ch := map[rune]byte{
		'⇖': '/',
		'⇗': '\\',
		'⇘': '/',
		'⇙': '\\',
		'⇐': '*',
		'⇒': '*',
		'⇑': '*',
		'⇓': '*',
		'🅂': ' ',
	}
	return Quad(x, y, ch)
}

func QuadC(x, y int) string {
	ch := map[rune]byte{
		'⇖': 'A',
		'⇗': 'A',
		'⇘': 'C',
		'⇙': 'C',
		'⇐': 'B',
		'⇒': 'B',
		'⇑': 'B',
		'⇓': 'B',
		'🅂': ' ',
	}
	return Quad(x, y, ch)
}

func QuadD(x, y int) string {
	ch := map[rune]byte{
		'⇖': 'A',
		'⇗': 'C',
		'⇘': 'C',
		'⇙': 'A',
		'⇐': 'B',
		'⇒': 'B',
		'⇑': 'B',
		'⇓': 'B',
		'🅂': ' ',
	}
	return Quad(x, y, ch)
}

func QuadE(x, y int) string {
	ch := map[rune]byte{
		'⇖': 'A',
		'⇗': 'C',
		'⇘': 'A',
		'⇙': 'C',
		'⇐': 'B',
		'⇒': 'B',
		'⇑': 'B',
		'⇓': 'B',
		'🅂': ' ',
	}
	return Quad(x, y, ch)
}

func main() {
	b := make([]byte, 1024)
	var quadToCheck string
	for {
		n, err := os.Stdin.Read(b[:])
		if n > 0 {
			quadToCheck += string(b[:n])
		}
		if err != nil {
			break
		}
	}
	if quadToCheck == "" {
		fmt.Print("Not a quad function\n")
		return
	}
	newLineCounter := 0
	length := 0 //?
	counterX := 0
	for _, value := range quadToCheck {
		if value == '\n' {
			if newLineCounter == 0 {
				length = counterX
			} else if counterX != length {
				fmt.Print("Not a quad function\n")
				return
			}
			counterX = 0
			newLineCounter++
		} else {
			counterX++
		}
	}
	if length == 0 || newLineCounter == 0 {
		fmt.Print("Not a quad function\n")
		return
	}
	funcArray := []func(int, int) string{QuadA, QuadB, QuadC, QuadD, QuadE}
	answer := ""
	added := false
	for i := 0; i < len(funcArray); i++ {
		if funcArray[i](length, newLineCounter) == quadToCheck {
			if !added {
				answer += "[quad" + string('A'+i) + "] [" + intToString(length) + "] [" + intToString(newLineCounter) + "]"
				added = true
			} else {
				answer += " || [quad" + string('A'+i) + "] [" + intToString(length) + "] [" + intToString(newLineCounter) + "]"
			}
		}
	}
	if answer == "" {
		fmt.Print("Not a quad function\n")
		return
	}
	fmt.Println(answer)
}
