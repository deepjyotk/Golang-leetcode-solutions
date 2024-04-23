package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{"Positive numbers", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"Include zero", []int{0, 4, 3, 0}, 0, []int{0, 3}},
		{"Negative numbers", []int{-3, 4, 3, 90}, 0, []int{0, 2}},
		{"No solution", []int{1, 2, 3}, 7, nil},
		{"Single element", []int{2}, 2, nil},
		{"Empty slice", []int{}, 0, nil},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := twoSum(tc.nums, tc.target)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("twoSum(%v, %d) = %v, want %v", tc.nums, tc.target, got, tc.want)
			}
		})
	}
}
