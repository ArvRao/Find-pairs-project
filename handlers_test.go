package main

import (
	"reflect"
	"testing"
)

func TestFindAllPairs(t *testing.T) {
	tests := []struct {
		numbers []int
		target  int
		want    [][]int
	}{
		// Test case 1: Standard case with multiple pairs
		{
			numbers: []int{1, 2, 3, 4, 5},
			target:  5,
			want:    [][]int{{0, 3}, {1, 2}},
		},
		// Test case 2: Case with duplicate numbers that form the target sum
		{
			numbers: []int{1, 2, 2, 3},
			target:  4,
			want:    [][]int{{0, 3}, {1, 2}},
		},
		// Test case 3: Case with no pairs summing to target
		{
			numbers: []int{1, 2, 3, 9},
			target:  20,
			want:    [][]int{},
		},
		// Test case 4: Case where all elements are zeroes, and target is zero
		{
			numbers: []int{0, 0, 0, 0},
			target:  0,
			want:    [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}, {2, 3}},
		},
		// Test case 5: Single number in input (no pairs possible)
		{
			numbers: []int{5},
			target:  5,
			want:    [][]int{},
		},
		// Test case 6: Empty input array
		{
			numbers: []int{},
			target:  5,
			want:    [][]int{},
		},
		// Test case 7: Large target with no pairs summing up
		{
			numbers: []int{10, 20, 30, 40, 50},
			target:  100,
			want:    [][]int{},
		},
		// Test case 8: Large input array where multiple pairs sum to the target
		{
			numbers: []int{1, 9, 2, 8, 3, 7, 4, 6, 5},
			target:  10,
			want:    [][]int{{0, 1}, {2, 3}, {4, 5}, {6, 7}},
		},
	}

	for _, tt := range tests {
		got := findAllPairs(tt.numbers, tt.target)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("findAllPairs(%v, %d) = %v; want %v", tt.numbers, tt.target, got, tt.want)
		}
	}
}
