package Array

import "fmt"

func MergeSortedArrayTest() {
	// test case 0
	//nums1 := []int{1, 2, 3, 0, 0, 0}
	//m := 3
	//nums2 := []int{2, 5, 6}
	//n := 3
	//merge(nums1, m, nums2, n)
	//fmt.Println(nums1)

	// test case 1
	//nums11 := []int{1}
	//m1 := 1
	//var nums21 []int
	//n1 := 0
	//merge(nums11, m1, nums21, n1)
	//fmt.Println(nums11)

	// test case 2
	nums12 := []int{0}
	m2 := 0
	nums22 := []int{1}
	n2 := 1
	merge(nums12, m2, nums22, n2)
	fmt.Println(nums12)
}

// 执行用时分布 击败 41.84% 使用 Go 的用户
// 消耗内存分布 击败 23.81% 使用 Go 的用户
//func merge(nums1 []int, m int, nums2 []int, n int) {
//	for i := 0; i < len(nums1); i++ {
//		if m == 0 || n == 0 {
//			break
//		}
//		if nums1[m-1] > nums2[n-1] {
//			nums1[m+n-1] = nums1[m-1]
//			m = m - 1
//		} else {
//			nums1[m+n-1] = nums2[n-1]
//			n = n - 1
//		}
//	}
//
//	if n > 0 && m == 0 {
//		for i := 0; i < n; i++ {
//			nums1[i] = nums2[i]
//		}
//	}
//}

// 执行用时分布 击败 41.84% 使用 Go 的用户
// 消耗内存分布 击败 10.92% 使用 Go 的用户
//func merge(nums1 []int, m int, nums2 []int, n int) {
//	for n > 0 {
//		if m > 0 && nums1[m-1] > nums2[n-1] {
//			nums1[m+n-1] = nums1[m-1]
//			m--
//		} else {
//			nums1[m+n-1] = nums2[n-1]
//			n--
//		}
//	}
//}

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums0 := make([]int, 0, m+n-1)
	i := 0
	j := 0
	for {
		if i == m {
			nums0 = append(nums0, nums2[j:]...)
			break
		}
		if j == n {
			nums0 = append(nums0, nums1[i:]...)
			break
		}
		if nums1[i] >= nums2[j] {
			nums0 = append(nums0, nums2[j])
			j++
		} else {
			nums0 = append(nums0, nums1[i])
			i++
		}
	}

	copy(nums1, nums0)
}
