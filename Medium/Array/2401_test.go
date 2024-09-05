package Array

func longestNiceSubarray(nums []int) int {
	left := 0
	maxLength := 0
	currentOr := 0
	for right, num := range nums {
		for currentOr&nums[right] != 0 {
			currentOr ^= nums[left]
			left++
		}
		currentOr |= num
		maxLength = max(maxLength, right-left+1)
	}

	return maxLength
}
