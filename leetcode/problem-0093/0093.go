package problem

import (
	"strconv"
	"strings"
)

func isValid(s string) bool {
	num, _ := strconv.Atoi(s)
	if num > 255 {
		return false
	}

	if strings.HasPrefix(s, "0") && len(s) != 1 {
		return false
	}

	return true
}

func backtrack(s string, ret *[]string, cur []string) {
	if len(s) == 0 && len(cur) == 4 {
		*ret = append(*ret, strings.Join(cur, "."))
		return
	}

	if len(cur) > 4 {
		return
	}

	for i := 1; i <= 3 && i <= len(s); i++ {
		if isValid(s[0:i]) {
			backtrack(s[i:], ret, append(cur, s[0:i]))
		}
	}
}

func restoreIpAddresses(s string) []string {
	var ret []string
	backtrack(s, &ret, nil)

	return ret
}
