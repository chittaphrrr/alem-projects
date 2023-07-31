package main

import (
	"fmt"
	"os"
)

func checkForCorrectness(s string) bool {
	for _, letter := range s {
		if letter < '0' || letter > '9' {
			return false
		}
	}
	return true
}

func StringToInt(s string) int {
	answer := 0
	for _, letter := range s {
		answer = answer*10 + int(letter-'0')
	}
	return answer
}

func QuadE(x, y int) string {
	var answer string
	if x > 0 && y > 0 {
		for i := 0; i < y; i++ {
			for j := 0; j < x; j++ {
				if i == 0 && j == 0 {
					answer += "A"
				} else if i == 0 && j == x-1 {
					answer += "C"
				} else if i == y-1 && j == 0 {
					answer += "C"
				} else if i == y-1 && j == x-1 {
					answer += "A"
				} else if i == 0 || i == y-1 || j == 0 || j == x-1 {
					answer += "B"
				} else {
					answer += " "
				}
			}
			answer += "\n"
		}
	}
	return answer
}

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		return
	}
	if checkForCorrectness(args[0]) && checkForCorrectness(args[1]) {
		x := StringToInt(args[0])
		y := StringToInt(args[1])
		fmt.Print(QuadE(x, y))
	}
}
