package problem0003

// lengthOfLongestSubstring is used to find the length of the longest substring without repeating characters.
// Time complexity: O(n)
// Space complexity: O(n)
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	left, right, n := 0, 0, len(s)
	maxLength := 0
	seen := make(map[uint8]int)

	for right < n {
		if _, ok := seen[s[right]]; ok {
			left = max(left, seen[s[right]]+1)
		}

		seen[s[right]] = right
		maxLength = max(maxLength, right-left+1)
		right++
	}

	return maxLength
}
