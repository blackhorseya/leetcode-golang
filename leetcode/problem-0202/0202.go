package problem_0202

func isHappy(n int) bool {
	sum, mod := 0, 0
	for ; n != 0; n /= 10 {
		mod = n % 10
		sum += mod * mod
	}

	if sum < 10 && sum == mod*mod {
		return sum == 1
	} else {
		return isHappy(sum)
	}
}
