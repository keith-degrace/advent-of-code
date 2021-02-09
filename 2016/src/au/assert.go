package au

import (
	"fmt"
)

func Assert(cond bool) {
	if !cond {
		panic("Assertion failed")
	}
}

func AssertStringsEqual(a string, b string) {
	if a != b {
		panic(fmt.Sprintf("Equality Assertion failed (%v != %v)", a, b))
	}
}

func AssertIntsEqual(a int, b int) {
	if a != b {
		panic(fmt.Sprintf("Equality Assertion failed (%v != %v)", a, b))
	}
}