package Array

import "fmt"

func RemoveDuplicatesTest() {
	// test case 0
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	numsLen := removeDuplicates(nums)
	fmt.Println(numsLen)
}

func removeDuplicates(nums []int) int {
	i := 0
	j := 0
	for i < len(nums) {
		if nums[i] != nums[j] {
			j++
			nums[j] = nums[i]
		}
		i++
	}
	return j + 1
}
