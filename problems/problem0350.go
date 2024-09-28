package problems

import (
	"sort"
)

// intersect 交集運算
// Time complexity: O(n+m)
// Space complexity: O(min(n,m))
func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) < len(nums2) { // 確保 nums1 是較長的 slice
		nums1, nums2 = nums2, nums1
	}

	// 執行交集運算
	// 建立 nums2 中每個元素的出現次數 [num]count
	m := make(map[int]int)
	for _, num := range nums2 {
		m[num]++
	}

	res := []int{}
	for _, num := range nums1 {
		count, ok := m[num]
		if !ok {
			continue
		} else {
			res = append(res, num)
			if count > 1 {
				m[num]--
			} else {
				delete(m, num)
			}
		}
	}

	return res
}

// intersect2 交集運算
// Time complexity: O(nlogn + mlogm)
// Space complexity: O(1)
func intersect2(nums1 []int, nums2 []int) []int {
	if len(nums1) < len(nums2) { // 確保 nums1 是較長的 slice
		nums1, nums2 = nums2, nums1
	}

	sort.Ints(nums1) // O(nlogn)
	sort.Ints(nums2) // O(mlogm)

	i, j := 0, 0
	res := []int{}
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			res = append(res, nums1[i])
			i++
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}

	return res
}
