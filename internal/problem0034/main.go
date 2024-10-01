package problem0034

// searchRange
func searchRange(nums []int, target int) []int {
	return []int{
		binarySearchLeftEdge(nums, target),
		binarySearchRightEdge(nums, target),
	}
}

/* 二分搜尋最左一個 target */
func binarySearchLeftEdge(nums []int, target int) int {
	// 等價於查詢 target 的插入點
	i := binarySearchInsertion(nums, target)
	// 未找到 target ，返回 -1
	if i == len(nums) || nums[i] != target {
		return -1
	}
	// 找到 target ，返回索引 i
	return i
}

/* 二分搜尋最右一個 target */
func binarySearchRightEdge(nums []int, target int) int {
	// 轉化為查詢最左一個 target + 1
	i := binarySearchInsertion(nums, target+1)
	// j 指向最右一個 target ，i 指向首個大於 target 的元素
	j := i - 1
	// 未找到 target ，返回 -1
	if j == -1 || nums[j] != target {
		return -1
	}
	// 找到 target ，返回索引 j
	return j
}

/* 二分搜尋插入點（存在重複元素） */
func binarySearchInsertion(nums []int, target int) int {
	// 初始化雙閉區間 [0, n-1]
	i, j := 0, len(nums)-1
	for i <= j {
		// 計算中點索引 m
		m := i + (j-i)/2
		if nums[m] < target {
			// target 在區間 [m+1, j] 中
			i = m + 1
		} else if nums[m] > target {
			// target 在區間 [i, m-1] 中
			j = m - 1
		} else {
			// 首個小於 target 的元素在區間 [i, m-1] 中
			j = m - 1
		}
	}
	// 返回插入點 i
	return i
}
