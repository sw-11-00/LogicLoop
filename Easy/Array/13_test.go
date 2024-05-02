package Array

import (
	"testing"
)

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "Test1",
			s:    "III",
			want: 3,
		},
		{
			name: "Test2",
			s:    "IV",
			want: 4,
		},
		{
			name: "Test3",
			s:    "IX",
			want: 9,
		},
		{
			name: "Test4",
			s:    "LVIII",
			want: 58,
		},
		{
			name: "Test5",
			s:    "MCMXCIV",
			want: 1994,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
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
