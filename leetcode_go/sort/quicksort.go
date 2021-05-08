package main

import (
	"fmt"
)

func main() {
	var arr = []int{15, 3, 12, 6, -9, 9, 0}

	QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)

}

func partition(arr []int, front, end int) int {
	pivot := arr[end]
	i := front - 1

	for j := front; j < end; j++ {
		if pivot > arr[j] {
			i++
			arr[j], arr[i] = arr[i], arr[j]
		}
	}
	i++
	arr[i], arr[end] = arr[end], arr[i]
	return i
}

func QuickSort(arr []int, front, end int) {
	fmt.Println(front, end)
	if front < end {

		pivot := partition(arr, front, end)
		QuickSort(arr, front, pivot-1)
		QuickSort(arr, pivot+1, end)
	}

}
