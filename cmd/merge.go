package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var mergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "Sorts a list of numbers using the merge sort algorithm",
	Long: `Merge sort is a divide-and-conquer algorithm that divides the list into
smaller sublists, sorts those sublists, and then merges them back together.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		nums, err := parseInput(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}

		if len(nums) < 2 {
			fmt.Println("List is already sorted:", nums)
			return
		}

		start := time.Now()

		mid := len(nums) / 2
		left := mergeSort(nums[:mid])
		right := mergeSort(nums[mid:])

		fmt.Println("Sorted list:", merge(left, right, nums))

		elapsed := time.Since(start)
		fmt.Printf("\nTotal time: %s\n", elapsed)
	},
}

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	mid := len(arr) / 2
	left := arr[:mid]
	right := arr[mid:]

	fmt.Printf("Dividing original list: %v into\nLeft list: %v\nRight list: %v\n", arr, left, right)

	leftList := mergeSort(left)
	rightList := mergeSort(right)
	merged := merge(leftList, rightList, arr)

	return merged
}

func merge(left, right, original []int) []int {
	sortedList := []int{}
	i := 0
	j := 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			sortedList = append(sortedList, left[i])
			i++
		} else {
			sortedList = append(sortedList, right[j])
			j++
		}
	}

	for ; i < len(left); i++ {
		sortedList = append(sortedList, left[i])
	}

	for ; j < len(right); j++ {
		sortedList = append(sortedList, right[j])
	}

	fmt.Printf("Original sub-list: %v; Sorted sub-list: %v\n", original, sortedList)

	return sortedList
}
