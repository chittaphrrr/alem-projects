package main

import (
	"goV2/functions"
	"log"
	"os"
	"strings"
)

func main() {
	if ok, err := functions.ValidateArguments(os.Args); !ok {
		log.Println("ERROR: ", err)
		return
	}

	content, err := os.ReadFile(os.Args[1])
	if strings.TrimSpace(string(content)) == "" {
		log.Println("ERROR: empty input file")
		return
	}
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}
	f, err := os.Create(os.Args[2])
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}

	if _, err = f.WriteString(functions.Convert(content)); err != nil {
		log.Println("ERROR: ", err)
		return
	}
}
