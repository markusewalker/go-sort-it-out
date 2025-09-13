package cmd

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
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
		result := quickSort(arr, 0, len(arr)-1)

		if !reflect.DeepEqual(result, c.expected) {
			t.Errorf("quickSort(%v) == %v, want %v", c.input, result, c.expected)
		}
	}
}
