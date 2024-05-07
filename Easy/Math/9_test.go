package Math

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want bool
	}{
		{
			name: "Test1",
			num:  101,
			want: true,
		},
		{
			name: "Test2",
			num:  -101,
			want: false,
		},
		{
			name: "Test3",
			num:  10,
			want: false,
		},
		{
			name: "Test4",
			num:  0,
			want: true,
		},
		{
			name: "Test5",
			num:  12321,
			want: true,
		},
		{
			name: "Test6",
			num:  -12321,
			want: false,
		},
		{
			name: "Test7",
			num:  11,
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPalindrome(tt.num)
			if got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	digits := digitToArray(x)
	for i, j := 0, len(digits)-1; i < j; {
		if digits[i] != digits[j] {
			return false
		}
		i++
		j--
	}

	return true
}

func digitToArray(x int) []int {
	var ans []int
	i := 0
	for x > 0 {
		ans = append(ans, x%10)
		x /= 10
		i++
	}

	return ans
}
