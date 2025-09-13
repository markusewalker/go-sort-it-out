package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var bubbleCmd = &cobra.Command{
	Use:   "bubble",
	Short: "Sorts a list of numbers using the bubble sort algorithm",
	Long: `Bubble sort is a simple sorting algorithm that repeatedly steps through the list,
compares adjacent elements and swaps them if they are in the wrong order.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		nums, err := parseInput(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}

		start := time.Now()

		bubbleSort(nums)

		elapsed := time.Since(start)
		fmt.Printf("\nTotal time: %s\n", elapsed)
	},
}

// parseInput parses a comma-separated string of numbers to be sorted.
func parseInput(input string) ([]int, error) {
	nums := []int{}
	split := strings.SplitSeq(input, ",")

	for s := range split {
		n, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil {
			return nil, fmt.Errorf("invalid number: %s", s)
		}

		nums = append(nums, n)
	}

	return nums, nil
}

func bubbleSort(arr []int) {
	fmt.Println("Unsorted list:", arr)

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}

		fmt.Printf("List after pass %d: %v\n", i+1, arr)
	}

	fmt.Println("Sorted list:", arr)
}
