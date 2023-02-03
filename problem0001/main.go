package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for flag, num := range nums {
		need := target - num
		ami, ok := m[need]
		if ok {
			return []int{ami, flag}
		}

		m[num] = flag
	}

	return []int{}
}

func main() {

}
