package Array

func maxArrayValue(nums []int) int64 {
	maxVal := nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] <= maxVal {
			maxVal += nums[i]
		} else {
			maxVal = nums[i]
		}
	}

	return int64(maxVal)
}
