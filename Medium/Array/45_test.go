package Array

func jump(nums []int) int {
	end := 0
	maxPosition := 0
	step := 0

	for i := 0; i < len(nums)-1; i++ {
		if i+nums[i] > maxPosition {
			maxPosition = i + nums[i]
		}

		if i == end {
			end = maxPosition
			step++
		}
	}

	return step
}
