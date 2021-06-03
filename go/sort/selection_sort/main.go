package main

import (
	"fmt"
)

func SelectionSort(arr []int) []int {
	var minIndex int
	length := len(arr)
	for i := 0; i < length-1; i++ {
		minIndex = i
		for j := i + 1; j < length; j++ {
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}
		temp := arr[i]
		arr[i] = arr[minIndex]
		arr[minIndex] = temp
	}
	return arr
}

func main() {
	// 13445789
	arr := []int{1, 8, 3, 9, 4, 5, 4, 7}
	SelectionSort(arr)
	for _, v := range arr {
		fmt.Print(v)
	}
}
