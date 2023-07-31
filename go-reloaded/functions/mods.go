package functions

import (
	"log"
	"strconv"
	"strings"
)

func DoMods(slice []string) []string {
	for i := 0; i < len(slice); i++ {
		switch {
		case slice[i] == "(up)":
			for k := i - 1; k >= 0; k-- {
				if slice[k] == "\n" || hasOnlyPunctuation(slice[k]) {
					continue
				} else {
					slice[k] = strings.ToUpper(slice[k])
					break
				}
			}
			slice[i] = ""
			i = 0
			slice = cleaner(slice)
		case slice[i] == "(low)":
			for k := i - 1; k >= 0; k-- {
				if slice[k] == "\n" || hasOnlyPunctuation(slice[k]) {
					continue
				} else {
					slice[k] = strings.ToLower(slice[k])
					break
				}
			}
			slice[i] = ""
			i = 0
			slice = cleaner(slice)
		case slice[i] == "(cap)":
			for k := i - 1; k >= 0; k-- {
				if slice[k] == "\n" || hasOnlyPunctuation(slice[k]) {
					continue
				} else {
					slice[k] = strings.Title(strings.ToLower(slice[k]))
					break
				}
			}
			slice[i] = ""
			i = 0
			slice = cleaner(slice)
		case slice[i] == "(up,":
			if i < len(slice)-1 {
				num, err := strconv.Atoi(slice[i+1][:len(slice[i+1])-1])
				if len(slice[i+1]) > 18 {
					log.Println("ERROR: overflow in ", slice[i]+" "+slice[i+1], " mod")
					slice[i], slice[i+1] = "", ""
					i = 0
					slice = cleaner(slice)
					continue
				}
				if err != nil || num < 0 {
					log.Println("ERROR: wrong number in: ", slice[i]+" "+slice[i+1])
				}
				if num > i {
					num = i
				}
				if num == 0 {
					num = -1
				}
				// fmt.Println("up index: ", i, "num: ", num)
				for k := i - 1; k >= i-num; k-- {
					if k < 0 {
						break
					}
					if slice[k] == "\n" || hasOnlyPunctuation(slice[k]) || slice[k] == "\\n" {
						num++
						continue
					}
					slice[k] = strings.ToUpper(slice[k])
					// fmt.Println(slice[k])
				}
				slice[i], slice[i+1] = "", ""
				i = 0
				slice = cleaner(slice)
			}
		case slice[i] == "(low,":
			if i < len(slice)-1 {
				num, err := strconv.Atoi(slice[i+1][:len(slice[i+1])-1])
				if len(slice[i+1]) > 18 {
					log.Println("ERROR: overflow in ", slice[i]+" "+slice[i+1], " mod")
					slice[i], slice[i+1] = "", ""
					i = 0
					slice = cleaner(slice)
					continue
				}
				if err != nil || num < 0 {
					log.Println("ERROR: wrong number in: ", slice[i]+" "+slice[i+1])
				}
				if num > i {
					num = i
				}
				if num == 0 {
					num = -1
				}
				// fmt.Println("low index: ", i, "num: ", num)
				for k := i - 1; k >= i-num; k-- {
					if k < 0 {
						break
					}
					if slice[k] == "\n" || hasOnlyPunctuation(slice[k]) || slice[k] == "\\n" {
						num++
						continue
					}
					slice[k] = strings.ToLower(slice[k])
					// fmt.Println(slice[k])
				}
				slice[i], slice[i+1] = "", ""
				i = 0
				slice = cleaner(slice)
			}
		case slice[i] == "(cap,":
			if i < len(slice)-1 {
				if len(slice[i+1]) > 18 {
					log.Println("ERROR: overflow in ", slice[i]+" "+slice[i+1], " mod")
					slice[i], slice[i+1] = "", ""
					i = 0
					slice = cleaner(slice)
					continue
				}
				num, err := strconv.Atoi(slice[i+1][:len(slice[i+1])-1])
				if err != nil || num < 0 {
					log.Println("ERROR: wrong number in: ", slice[i]+" "+slice[i+1])
				}
				if num > i {
					num = i
				}
				if num == 0 {
					num = -1
				}
				// fmt.Println("cap index: ", i, "num: ", num)
				for k := i - 1; k >= i-num; k-- {
					if k < 0 {
						break
					}
					if slice[k] == "\n" || hasOnlyPunctuation(slice[k]) || slice[k] == "\\n" {
						num++
						continue
					}
					slice[k] = strings.Title(strings.ToLower(slice[k]))
					// fmt.Println(slice[k])
				}
				slice[i], slice[i+1] = "", ""
				i = 0
				slice = cleaner(slice)
			}
		case slice[i] == "(bin)":
			if i > 0 && isBin(slice[i-1]) {
				slice[i-1] = convertBin(slice[i-1])
			} else {
				log.Println("ERROR: non-binary number for (bin) in ", slice[i-1], " word")
			}
			slice[i] = ""
			i = 0
			slice = cleaner(slice)
		case slice[i] == "(hex)":
			if i > 0 && isHex(slice[i-1]) {
				slice[i-1] = convertHex(slice[i-1])
			} else {
				log.Println("ERROR: non-hexidecimal number for (hex) in ", slice[i-1], " word")
			}
			slice[i] = ""
			i = 0
			slice = cleaner(slice)
		}
	}

	return cleaner(slice)
}

func convertBin(s string) string {
	temp, _ := strconv.ParseInt(s, 2, 64)
	s = strconv.Itoa(int(temp))
	return s
}

func isBin(s string) bool {
	if s == "\n" {
		return false
	}
	_, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		return false
	}
	return true
}

func convertHex(s string) string {
	temp, _ := strconv.ParseInt(s, 16, 64)
	s = strconv.Itoa(int(temp))
	return s
}

func isHex(s string) bool {
	if s == "\n" {
		return false
	}
	_, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		return false
	}
	return true
}
