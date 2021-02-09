package p14_1

import (
	"fmt"
	"io"
	"crypto/md5"
)

func getHash(value string) string {
	hash := md5.New()
	io.WriteString(hash, value)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func getFirstTriplet(hash string) string {
	for i := 0; i < len(hash) - 2; i++ {
		if hash[i] == hash[i+1] && hash[i] == hash[i+2] {
			return string(hash[i])
		}
	}

	return ""
}

func hasQuintuplet(hash string, tripletValue string) bool {
	runCount := 0

	for i := 0; i < len(hash); i++ {
		if string(hash[i]) == tripletValue {
			runCount++
			if runCount == 5 {
				return true
			}
		} else {
			runCount = 0
		}
	}

	return false
}

func hasFollowingQuintuplet(startingIndex int, tripletValue string, salt string, hashes *map[string] string) bool {
	for i := 0; i < 1000; i++ {
		saltedValue := fmt.Sprintf("%v%v", salt, startingIndex + i) 
		hash,ok := (*hashes)[saltedValue]
		if !ok {
			hash = getHash(saltedValue)
			(*hashes)[saltedValue] = hash
		}

		if hasQuintuplet(hash, tripletValue) {
			return true
		}
	}

	return false
}

func Solve() {
	salt := "ahsbgdzn"
	// salt = "abc"

	hashes := map[string] string{}

	index := 0
	keyCount := 0
	for {
		saltedValue := fmt.Sprintf("%v%v", salt, index)
		hash,ok := hashes[saltedValue]
		if !ok {
			hash = getHash(saltedValue)
		}

		tripletValue := getFirstTriplet(hash)
		if tripletValue != "" {
			if hasFollowingQuintuplet(index + 1, tripletValue, salt, &hashes) {
				keyCount++
				if keyCount == 64 {
					fmt.Println(index)
					return
				}
			}	
		}

		index++
	}
}

