package p07_1

import (
	"au"
	"fmt"
	"strings"
)

func testInputs() []string {
	return []string {
		"abba[mnop]qrst",
		"abcd[bddb]xyyx",
		"aaaa[qwer]tyui",
		"ioxxoj[asdfgh]zxcvbn",
	};
}

func isABBA(part string) bool {
	for i := 0; i < len(part) - 3; i++ {
		if (part[i] != part[i+1] &&
			  part[i] == part[i+3] &&
			  part[i+1] == part[i+2]) {
			return true;
		}
	}

	return false;
}

func hasABBA(parts []string) bool {
	for _, part := range parts {
		if isABBA(part) {
			return true;
		}
	}

	return false
}

func supportsTLS(ip string) bool {
	parts := strings.Split(strings.Replace(ip, "[", "]", -1), "]")

	supernetParts := []string {}
	hypernetParts := []string {}

	for i := 0; i < len(parts); i++ {
		if i % 2 == 0 {
			supernetParts = append(supernetParts, parts[i])
		} else {
			hypernetParts = append(hypernetParts, parts[i])
		}
	}

	return hasABBA(supernetParts) && !hasABBA(hypernetParts)
}

func Solve() {
	ips := au.ReadInputAsStringArray("07")
	// ips := testInputs()

	count := 0
	for _, ip := range ips {
		if (supportsTLS(ip)) {
			count++;
		}
	}

	fmt.Println(count)
}
