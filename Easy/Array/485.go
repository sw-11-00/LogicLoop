package Array

func findMaxConsecutiveOnes(nums []int) int {
	maxConsecutiveOnes := 0
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			count = 0
		} else {
			count++
		}
		maxConsecutiveOnes = max(maxConsecutiveOnes, count)
	}

	return maxConsecutiveOnes
}
