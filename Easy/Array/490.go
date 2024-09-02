package Array

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m, n := len(nums1), len(nums2)
	res := make([]int, m)
	for i, num1 := range nums1 {
		j := 0
		for j < n && nums2[j] != num1 {
			j++
		}
		j = j + 1
		for j < n && nums2[j] <= num1 {
			j++
		}
		if j == n {
			res[i] = -1
		} else {
			res[i] = nums2[j]
		}
	}

	return res
}
