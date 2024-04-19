package Array

import (
	"fmt"
)

func RomanToIntTest() {
	testCases := []struct {
		roman    string
		expected int
	}{
		{"III", 3},
		{"IV", 4},
		{"IX", 9},
		{"LVIII", 58},     // Explanation: L = 50, V= 5, III = 3.
		{"MCMXCIV", 1994}, // Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
		{"MMXX", 2020},
		{"CDXLIV", 444},
	}

	for _, tc := range testCases {
		result := romanToInt(tc.roman)
		if result == tc.expected {
			fmt.Printf("Test Passed for '%s': Expected and got %d\n", tc.roman, result)
		} else {
			fmt.Printf("Test Failed for '%s': Expected %d, Got %d\n", tc.roman, tc.expected, result)
		}
	}
}

func romanToInt(s string) int {
	valueMap := getRomanValueMap()
	sum := 0
	for i := len(s) - 1; i >= 0; i-- {
		if i-1 >= 0 && valueMap[s[i]] > valueMap[s[i-1]] {
			sum += valueMap[s[i]] - 2*valueMap[s[i-1]]
		} else {
			sum += valueMap[s[i]]
		}
	}
	return sum
}

func getRomanValueMap() map[byte]int {
	return map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
}
