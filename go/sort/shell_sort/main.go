package main

import (
	"fmt"
)

func swap(arr []int, a, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}

func ShellSort(arr []int) []int {
	length := len(arr)
	gap := 2
	step := length / gap
	for step >= 1 {
		for i := step; i < length; i++ {
			for j := i; j >= step && arr[j] < arr[j-step]; j -= step {
				swap(arr, j, j-step)
			}
		}
		step /= gap
	}
	return arr
}

func main() {
	// 13445789
	arr := []int{1, 8, 3, 9, 4, 5, 4, 7}
	ShellSort(arr)
	for _, v := range arr {
		fmt.Print(v)
	}
}
