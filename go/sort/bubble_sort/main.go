package main

import "fmt"

func BubbleSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if arr[j] > arr[j+1] {
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}
	return arr
}

func main() {
	// 13445789
	arr := []int{1, 8, 3, 9, 4, 5, 4, 7}
	BubbleSort(arr)
	for _, v := range arr {
		fmt.Print(v)
	}
}
