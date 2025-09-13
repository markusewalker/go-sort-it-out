package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var insertionCmd = &cobra.Command{
	Use:   "insertion",
	Short: "Sorts a list of numbers using the insertion sort algorithm",
	Long: `Insertion sort is a simple sorting algorithm that builds the final sorted array 
	one item at a time. It is much less efficient on larger lists.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		nums, err := parseInput(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}

		start := time.Now()

		insertionSort(nums)

		elapsed := time.Since(start)
		fmt.Printf("\nTotal time: %s\n", elapsed)
	},
}

func insertionSort(arr []int) {
	fmt.Println("Unsorted list:", arr)

	for i := range arr {
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}

		fmt.Printf("List after pass %d: %v\n", i+1, arr)
	}

	fmt.Println("Sorted list:", arr)
}
