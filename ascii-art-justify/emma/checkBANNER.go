package emma

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io/ioutil"
)

func HashMD5(content string) string {
	h := md5.Sum([]byte(content))
	return fmt.Sprintf("%x", h)
}

func CheckHash(format string) error {
	format = "./data/" + format + ".txt"
	arrByte, err := ioutil.ReadFile(format)
	if err != nil {
		return err
	}
	hash := HashMD5(string(arrByte))
	// fmt.Println(hash)
	hashStandard := "ac85e83127e49ec42487f272d9b9db8b"
	hashShadow := "a49d5fcb0d5c59b2e77674aa3ab8bbb1"
	hashThinkertoy := "86d9947457f6a41a18cb98427e314ff8"
	if hash == hashStandard {
		return nil
	}
	if hash == hashShadow {
		return nil
	}
	if hash == hashThinkertoy {
		return nil
	}
	return errors.New("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard\nYour BANNER  has been changed")
}
