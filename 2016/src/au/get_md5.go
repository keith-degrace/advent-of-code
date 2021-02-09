package au

import (
	"crypto/md5"
	"fmt"
	"io"
)

func GetMD5(value string) string {
	hash := md5.New()
	io.WriteString(hash, value)
	return fmt.Sprintf("%x", hash.Sum(nil))
}
