# Go-Sort-It-Out

| Tests |
|:-----:|
| [![Tests](https://github.com/markusewalker/go-sort-it-out/actions/workflows/run-tests.yml/badge.svg?branch=main)](https://github.com/markusewalker/go-sort-it-out/actions/workflows/run-tests.yml) |


Go-Sort-It-Out is a basic sorting algorithm CLI built using Cobra. As of now, it supports the following sorting algorithms:

- Bubble sort
- Insertion sort
- Merge sort
- Quick sort

## Table of Contents
1. [Prerequisites](#Prerequisites)
2. [Getting Started](#Getting-Started)
3. [Testing](#Testing)

## Prerequisites
In order to properly run this, you will need the following tools installed and configured on your client machine:

- Go (1.25+)

## Getting Started
To get going, there are two options.
1. Directly running `go run main.go <sorting option> <comma-separated list of numbers>`
2. Downloading the relevant binary in the Releases pages

Either option is entirely up to you! Here is the usage of the CLI tool:

```
$ go run main.go 
Go Sort It Out is a CLI application that allows users to sort a list of numbers
using various sorting algorithms. Users can choose the algorithm they want to use and input their list 
of numbers to see the sorted output.

Usage:
  gosort [command]

Available Commands:
  bubble      Sorts a list of numbers using the bubble sort algorithm
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  insertion   Sorts a list of numbers using the insertion sort algorithm
  merge       Sorts a list of numbers using the merge sort algorithm
  quick       Sorts a list of numbers using the quick sort algorithm

Flags:
  -h, --help   help for gosort

Use "gosort [command] --help" for more information about a command.
```