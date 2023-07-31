package emma

import (
	"fmt"
	"strings"
)

func Thor(ascistr, args string, qatar int) {
	ascistr = strings.ReplaceAll(ascistr, "\r", "")
	ascimas := strings.Split(ascistr, "\n")
	str := ""
	for _, elem := range []rune(args) {
		str += ascimas[(int(elem)-32)*9+qatar]
	}
	fmt.Println(str)
}

func ThorR(ascistr, args string, qatar int) bool {
	ascistr = strings.ReplaceAll(ascistr, "\r", "")
	ascimas := strings.Split(ascistr, "\n")
	str := ""
	for _, elem := range []rune(args) {
		str += ascimas[(int(elem)-32)*9+qatar]
	}
	weight, _, _ := GetTerminalSize()
	if len(str) > weight {
		fmt.Println("Your text does not fit on the terminal window\nUsage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
		return false
	}
	fmt.Println(str)
	return true
}

func ThorL(ascistr, args string, qatar int) bool {
	ascistr = strings.ReplaceAll(ascistr, "\r", "")
	ascimas := strings.Split(ascistr, "\n")
	str := ""
	for _, elem := range []rune(args) {
		str += ascimas[(int(elem)-32)*9+qatar]
	}
	weight, _, _ := GetTerminalSize()
	if len(str) > weight {
		fmt.Println("Your text does not fit on the terminal window\nUsage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
		return false
	} else {
		probels := ""
		probelsnum := (weight - len(str))
		for i := 0; i < probelsnum; i++ {
			probels += " "
		}
		fmt.Println(probels + str)
	}

	return true
}

func ThorC(ascistr, args string, qatar int) bool {
	ascistr = strings.ReplaceAll(ascistr, "\r", "")
	ascimas := strings.Split(ascistr, "\n")
	str := ""
	for _, elem := range []rune(args) {
		str += ascimas[(int(elem)-32)*9+qatar]
	}
	weight, _, _ := GetTerminalSize()
	if len(str) > weight {
		fmt.Println("Your text does not fit on the terminal window\nUsage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
		return false
	} else {
		probels := ""
		probelsnum := (weight - len(str))
		if probelsnum%2 != 0 {
			probelsnum -= 1
		}
		for i := 0; i < probelsnum/2; i++ {
			probels += " "
		}
		fmt.Println(probels + str + probels)
	}

	return true
}

func ThorJ(ascistr, args string, qatar int) bool {
	weight, _, _ := GetTerminalSize()
	ascistr = strings.ReplaceAll(ascistr, "\r", "")
	ascimas := strings.Split(ascistr, "\n")
	str := ""
	for _, elem := range []rune(args) {
		str += ascimas[(int(elem)-32)*9+qatar]
	}
	if len(str) > weight {
		fmt.Println("Your text does not fit on the terminal window\nUsage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
		return false
	}
	cntspas := countSpaces(args)
	pust := weight - len(str) + cntspas*6
	args = removeExtraSpaces(args)
	argsmas := strings.Split(args, " ")
	probelsnum := len(argsmas) - 1

	if probelsnum == 0 || len(argsmas) == 1 {
		ThorC(ascistr, args, qatar)
		return true
	}
	interpustnum := pust / probelsnum
	interpust := ""
	for i := 0; i < interpustnum; i++ {
		interpust += " "
	}
	lishpustnum := pust % probelsnum
	lishpust := ""
	for i := 0; i < lishpustnum; i++ {
		lishpust += " "
	}
	str = ""
	check := true
	for h, soz := range argsmas {
		for _, elem := range []rune(soz) {
			str += ascimas[(int(elem)-32)*9+qatar]
		}
		if h != len(argsmas)-1 {
			str += interpust
		}
		if h == (len(argsmas)-1)/2 && check {
			str += lishpust
			check = false
		}
	}
	// str = ""
	// for _, elem := range []rune(args) {
	// 	str += ascimas[(int(elem)-32)*9+qatar]
	// }
	fmt.Println(str)
	// if qatar == 8 {
	// 	fmt.Println("args:", args, "  argsmas: ", argsmas, " probelsnum:", probelsnum, "interpust:", interpustnum, "lisjjgg:", lishpustnum)
	// }
	return true
}

func removeExtraSpaces(text string) string {
	words := strings.Fields(text)
	cleanText := strings.Join(words, " ")
	return cleanText
}

func countSpaces(text string) int {
	count := 0
	for _, char := range text {
		if char == ' ' {
			count++
		}
	}
	return count
}
