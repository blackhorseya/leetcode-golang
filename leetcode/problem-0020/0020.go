package problem

func isValid(s string) bool {
	var tmp []string

	for _, c := range s {
		switch c {
		case '(':
			tmp = append(tmp, string(c))
			break
		case '{':
			tmp = append(tmp, string(c))
			break
		case '[':
			tmp = append(tmp, string(c))
			break
		case ')':
			if len(tmp) == 0 {
				return false
			}
			if end := tmp[len(tmp)-1]; end != "(" {
				return false
			}
			tmp = append(tmp[:len(tmp)-1], tmp[len(tmp)-1+1:]...)
			break
		case ']':
			if len(tmp) == 0 {
				return false
			}
			if end := tmp[len(tmp)-1]; end != "[" {
				return false
			}
			tmp = append(tmp[:len(tmp)-1], tmp[len(tmp)-1+1:]...)
			break
		case '}':
			if len(tmp) == 0 {
				return false
			}
			if end := tmp[len(tmp)-1]; end != "{" {
				return false
			}
			tmp = append(tmp[:len(tmp)-1], tmp[len(tmp)-1+1:]...)
			break
		}
	}

	if len(tmp) != 0 {
		return false
	}

	return true
}
