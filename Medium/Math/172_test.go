package Math

import (
	"testing"
)

func TestTrailingZeroes(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "Test1",
			n:    3,
			want: 0,
		},
		{
			name: "Test2",
			n:    5,
			want: 1,
		},
		{
			name: "Test3",
			n:    0,
			want: 0,
		},
		{
			name: "Test4",
			n:    10,
			want: 2,
		},
		{
			name: "Test5",
			n:    25,
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trailingZeroes(tt.n); got != tt.want {
				t.Errorf("trailingZeroes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func trailingZeroes(n int) int {
	count := 0
	for n > 0 {
		n /= 5
		count += n
	}
	return count
}
