package Array

import (
	"fmt"
)

func LongestCommonPrefixTest() {
	testCases := []struct {
		strs     []string
		expected string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"interspecies", "interstellar", "interstate"}, "inters"},
		{[]string{"throne", "throne"}, "throne"},
		{[]string{}, ""},
		{[]string{"prefix", "prefix"}, "prefix"},
		{[]string{"", "b", "c"}, ""},
	}

	for _, tc := range testCases {
		result := longestCommonPrefix(tc.strs)
		if result == tc.expected {
			fmt.Printf("Test Passed for %v: Expected and got '%s'\n", tc.strs, result)
		} else {
			fmt.Printf("Test Failed for %v: Expected '%s', Got '%s'\n", tc.strs, tc.expected, result)
		}
	}
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	commonStr := comp(strs[0], strs[1])

	j := 2
	for j < len(strs) {
		newCommonStr := comp(commonStr, strs[j])
		if newCommonStr == "" {
			return newCommonStr
		}
		j++
		commonStr = newCommonStr
	}

	return commonStr
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func comp(a, b string) string {
	commonStr := ""
	strLen := min(len(a), len(b))
	for i := 0; i < strLen; i++ {
		if a[i] == b[i] {
			commonStr = commonStr + string(a[i])
		} else {
			return commonStr
		}
	}

	return commonStr
}
