package problems

func isValid(s string) bool {
	stack := make([]uint8, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' { // push to stack
			stack = append(stack, s[i])
		} else { // pop from stack
			if len(stack) == 0 {
				return false
			}

			peek := stack[len(stack)-1]
			switch s[i] {
			case ')':
				if peek != '(' {
					return false
				}
			case ']':
				if peek != '[' {
					return false
				}
			case '}':
				if peek != '{' {
					return false
				}
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
