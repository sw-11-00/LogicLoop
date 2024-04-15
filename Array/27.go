package Array

import "fmt"

func RemoveElementTest() {
	// test case 0
	nums := []int{3, 2, 2, 3}
	val := 3
	numsLen := removeElement(nums, val)
	fmt.Println(numsLen)
}

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	j := 0

	// 2
	for i < len(nums) {
		if nums[i] != val {
			nums[j] = nums[i]
			j++
		}
		i++
	}

	return j
}

func removeElement2(nums []int, val int) int {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}
