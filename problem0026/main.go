package main

func removeDuplicates(nums []int) int {
	k := 1

	for i := range nums {
		if i == 0 {
			continue
		}

		if nums[k-1] == nums[i] {
			continue
		} else {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}

func main() {

}
