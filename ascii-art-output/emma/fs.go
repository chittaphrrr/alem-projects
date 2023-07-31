package emma

import (
	"io/ioutil"
)

func Fs(arg string) string {
	var text string
	if arg == "shadow" {
		textb, _ := ioutil.ReadFile("data/shadow.txt")
		// if err != nil {
		// 	fmt.Println(err)
		// 	return "err"
		// }
		text = string(textb)
	} else if arg == "thinkertoy" {
		textb, _ := ioutil.ReadFile("data/thinkertoy.txt")
		text = string(textb)
	} else if arg == "standard" {
		textb, _ := ioutil.ReadFile("data/standard.txt")
		text = string(textb)
	}

	return text
}
