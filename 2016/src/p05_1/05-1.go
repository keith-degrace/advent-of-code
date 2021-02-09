package p05_1

import (
	//"au"
	"crypto/md5"
	"fmt"
	"io"
	"strings"
)

func getMD5(value string) string {
	hash := md5.New()
	io.WriteString(hash, value)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func Solve() {
	input := "ugkcyxxp"
	// input := "abc"

	index := 0

	password := ""
	for len(password) < 8 {
		md5 := getMD5(fmt.Sprintf("%v%v", input, index))

		if strings.HasPrefix(md5, "00000") {
			password += string(md5[5])
		}

		index++
	}

	fmt.Println(password)
}
