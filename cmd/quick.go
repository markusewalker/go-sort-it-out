package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// quickCmd represents the quick command
var quickCmd = &cobra.Command{
	Use:   "quick",
	Short: "Sorts a list of numbers using the quick sort algorithm",
	Long: `Quick sort is an in-place sorting algorithm that uses a divide-and-conquer approach
to sort elements. It selects a 'pivot' element from the array and partitions the other
elements into two sub-arrays according to whether they are less than or greater than the pivot.`,
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
		sorted := sortList(nums)
		elapsed := time.Since(start)

		fmt.Println("Sorted list:", sorted)
		fmt.Printf("\nTotal time: %s\n", elapsed)
	},
}

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[high] = arr[high], arr[i]

	return arr, i
}

func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int

		arr, p = partition(arr, low, high)
		fmt.Printf("List after partitioning (low=%d, high=%d, pivot index=%d): %v\n", low, high, p, arr)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}

	return arr
}

func sortList(arr []int) []int {
	fmt.Println("Unsorted list:", arr)

	return quickSort(arr, 0, len(arr)-1)
}
