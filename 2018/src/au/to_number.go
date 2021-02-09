package au

import (
	"strconv"
	"strings"
)

func ToNumber(s string) int {
	i, err := strconv.Atoi(strings.TrimSpace(s));
	FatalOnError(err);
	return i;
}
