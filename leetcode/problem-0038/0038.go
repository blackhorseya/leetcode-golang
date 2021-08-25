package problem

import (
	"strconv"
	"strings"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	var ret strings.Builder
	str := countAndSay(n - 1)
	count := 1
	for pos := 1; pos < len(str); pos++ {
		if str[pos] == str[pos-1] {
			count++
		} else {
			ret.WriteString(strconv.Itoa(count))
			ret.WriteByte(str[pos-1])
			count = 1
		}
	}

	ret.WriteString(strconv.Itoa(count))
	ret.WriteByte(str[len(str)-1])

	return ret.String()
}
