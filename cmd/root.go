package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gosort",
	Short: "A CLI tool to sort a list of numbers using different sorting algorithms",
	Long: `Go Sort It Out is a CLI application that allows users to sort a list of numbers
using various sorting algorithms. Users can choose the algorithm they want to use and input their list 
of numbers to see the sorted output.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(bubbleCmd)
	rootCmd.AddCommand(insertionCmd)
	rootCmd.AddCommand(mergeCmd)
	rootCmd.AddCommand(quickCmd)
}
