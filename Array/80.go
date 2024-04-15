package Array

import "fmt"

func RemoveDuplicates2Test() {
	// test case 0
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	numsLen := removeDuplicates2(nums)
	fmt.Println(numsLen)
}

func removeDuplicates2(nums []int) int {
	j := 0
	cnt := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[j] {
			nums[j] = nums[i]
			cnt = 0
		} else if nums[i] == nums[j] {
			if cnt == 2 {
				continue
			}
			j++
			cnt++
		}
	}

	return j
}
