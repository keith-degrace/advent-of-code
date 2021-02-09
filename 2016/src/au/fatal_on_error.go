package au

import (
	"fmt"
	"os"
)

func FatalOnError(err error) {
	if err != nil {
		fmt.Print(err)
		os.Exit(3)
	}
}
