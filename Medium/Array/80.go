package Array

import "fmt"

func RemoveDuplicates2Test() {
	// test case 0
	nums := []int{1, 1, 1, 2, 2, 3}
	numsLen := removeDuplicates2(nums)
	fmt.Println(numsLen)
}

func removeDuplicates2(nums []int) int {
	k := 0
	for _, num := range nums {
		if k < 2 || nums[k-2] < num {
			nums[k] = num
			k++
		}
	}

	return k
}
