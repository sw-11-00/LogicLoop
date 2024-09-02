package Array

func findDisappearedNumbers(nums []int) []int {
	numMap := make(map[int]bool)
	for _, num := range nums {
		numMap[num] = true
	}

	var res []int
	for i := 1; i <= len(nums); i++ {
		if !numMap[i] {
			res = append(res, i)
		}
	}

	return res
}
