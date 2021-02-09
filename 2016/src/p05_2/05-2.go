package p05_2

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
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

	password := []byte{'_', '_', '_', '_', '_', '_', '_', '_'}
	foundCount := 0
	for foundCount < 8 {
		md5 := getMD5(fmt.Sprintf("%v%v", input, index))

		if strings.HasPrefix(md5, "00000") {
			position, err := strconv.Atoi(string(md5[5]))
			if err == nil && position < 8 && password[position] == '_' {
				password[position] = md5[6]
				foundCount++
			}
		}

		index++
	}

	fmt.Println(string(password[:]))
}
