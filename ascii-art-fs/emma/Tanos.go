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
