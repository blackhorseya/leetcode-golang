package problem_0014

import "strings"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	for _, str := range strs {
		if len(str) == 0 {
			return ""
		}
	}

	pref := ""
	for i := 0; i < len(strs[0]); i++ {
		pref += string(strs[0][i])
		for j := 1; j < len(strs); j++ {
			if strings.HasPrefix(strs[j], pref) {
				continue
			} else {
				if i == 0 {
					return ""
				}

				return pref[:i]
			}
		}
	}

	return pref
}
