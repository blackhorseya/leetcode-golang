package problem_0088

func merge(nums1 []int, m int, nums2 []int, n int) {
	ret := make([]int, len(nums1))
	i, j, k := 0, 0, 0

	for i < m && j < n {
		if nums1[i] <= nums2[j] {
			ret[k] = nums1[i]
			i++
		} else {
			ret[k] = nums2[j]
			j++
		}

		k++
	}

	for i < m {
		ret[k] = nums1[i]
		i++
		k++
	}

	for j < n {
		ret[k] = nums2[j]
		j++
		k++
	}

	copy(nums1, ret)
}
