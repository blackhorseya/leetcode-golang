package problem

import "sort"

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) < len(nums2) {
		return intersect(nums2, nums1)
	}

	maps := make(map[int]int)
	for _, num := range nums1 {
		maps[num]++
	}

	var ret []int
	for _, num := range nums2 {
		count, ok := maps[num]
		if ok && count != 0 {
			ret = append(ret, num)
			maps[num]--
		}
	}

	sort.Ints(ret)
	return ret
}
