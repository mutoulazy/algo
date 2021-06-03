package main

import "fmt"

func InsertionSort(arr []int) []int {
	length := len(arr)
	var preIndex, current int
	for i := 1; i < length; i++ {
		preIndex = i - 1
		current = arr[i]
		for preIndex >= 0 && arr[preIndex] > current {
			arr[preIndex+1] = arr[preIndex]
			preIndex--
		}
		arr[preIndex+1] = current
	}
	return arr
}

func main() {
	// 13445789
	arr := []int{1, 8, 3, 9, 4, 5, 4, 7}
	InsertionSort(arr)
	for _, v := range arr {
		fmt.Print(v)
	}
}
