package main

import (
	"testing"
)

func TestTribonacci(t *testing.T) {
	testCases := []struct {
		name string
		n    int
		want int
	}{
		{"Lower Number1 ", 1, 1},
		{"Lower Number2", 0, 0},
		{"Other", 25, 1389537},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tribonacci(tc.n)
			if got != tc.want {
				t.Errorf("tribonacci(%d) = %d, want %d", tc.n, got, tc.want)
			}
		})
	}
}
