package emma

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
)

func HashMD5(content string) string {
	h := md5.Sum([]byte(content))
	return fmt.Sprintf("%x", h)
}

func CheckHash(format string) bool {
	format = "./data/" + format + ".txt"
	arrByte, err := ioutil.ReadFile(format)
	if err != nil {
		log.Println("Error! Incorrect read file from /data/" + format)
		return false
	}
	hash := HashMD5(string(arrByte))
	// fmt.Println(hash)
	hashStandard := "ac85e83127e49ec42487f272d9b9db8b"
	hashShadow := "a49d5fcb0d5c59b2e77674aa3ab8bbb1"
	// hashThinkertoy := "cd9ba1cc7a1b626147cf6729ad0c6857"
	hashThinkertoy := "86d9947457f6a41a18cb98427e314ff8"
	if hash == hashStandard {
		return true
	}
	if hash == hashShadow {
		return true
	}
	if hash == hashThinkertoy {
		return true
	}
	return false
}
