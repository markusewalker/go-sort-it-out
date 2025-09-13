package cmd

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	cases := []struct {
		input    []int
		expected []int
	}{
		{[]int{5, 2, 9, 1, 5, 6}, []int{1, 2, 5, 5, 6, 9}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
	}
	for _, c := range cases {
		arr := make([]int, len(c.input))
		copy(arr, c.input)
		insertionSort(arr)

		if !reflect.DeepEqual(arr, c.expected) {
			t.Errorf("insertionSort(%v) == %v, want %v", c.input, arr, c.expected)
		}
	}
}
