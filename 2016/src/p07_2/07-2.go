package p07_2

import (
	"au"
	"fmt"
	"strings"
)

func testInputs() []string {
	return []string {
		"aba[bab]xyz",
		"xyx[xyx]xyx",
		"aaa[kek]eke",
		"zazbz[bzb]cdb",
	};
}

func getABAs(part string) []string {
	abas := []string{}

	for i := 0; i < len(part) - 2; i++ {
		if (part[i] != part[i+1] && part[i] == part[i+2]) {
			abas = append(abas, string([]byte { part[i], part[i+1], part[i+2] }))
		}
	}

	return abas;
}

func supportsSSL(ip string) bool {
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

	supernetABAs := map [string] bool{}
	for _, supernetPart := range supernetParts {
		for _, aba := range getABAs(supernetPart) {
			supernetABAs[aba] = true;
		}
	}

	for _, hypernetPart := range hypernetParts {
		for _, aba := range getABAs(hypernetPart) {
			bab := string(aba[1]) + string(aba[0]) + string(aba[1])
			
			_, ok := supernetABAs[bab]
			if ok {
				return true;
			}
		}
	}

	return false;
}

func Solve() {
	ips := au.ReadInputAsStringArray("07")
	// ips := testInputs()

	count := 0
	for _, ip := range ips {
		if (supportsSSL(ip)) {
			count++;
		}
	}

	fmt.Println(count)
}
