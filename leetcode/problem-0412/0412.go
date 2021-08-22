package problem_0412

import "strconv"

func fizzBuzz(n int) []string {
	ret := make([]string, n)
	for i := 1; i <= n; i++ {
		var ch string
		if i%3 == 0 && i%5 == 0 {
			ch = "FizzBuzz"
		} else if i%5 == 0 {
			ch = "Buzz"
		} else if i%3 == 0 {
			ch = "Fizz"
		} else {
			ch = strconv.Itoa(i)
		}

		ret[i-1] = ch
	}

	return ret
}
