package main

import (
	"ascii-art/emma"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		log.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
	// 	}
	// }()
	var test string
	flag.StringVar(&test, "test", "DEFAULT", "u can align the text as u wish here")
	flag.Parse()
	fmt.Println(test)

	fmt.Println(len(os.Args))

	if len(os.Args) == 2 {
		err := emma.Asciiart()
		if err != nil {
			log.Fatal(err)
		}
		return
	} else if len(os.Args) == 3 {
		emma.Asciiartfs()
		return
	}
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
		return
	}
	emma.Asciiartjustify()
}

// [hello]\n[]\n[]\n[world]
