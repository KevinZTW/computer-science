package main

import "fmt"

func main() {
	fmt.Println("week 3")
	fmt.Println(MergeSort([]int{4, 2, 6, 10, 1, 0}))
}

func MergeSort(arr []int) []int {
	fmt.Println(arr)

	if len(arr) < 2 {
		fmt.Println("retunr")
		return arr
	}

	left := MergeSort(arr[:len(arr)/2])
	right := MergeSort(arr[len(arr)/2:])
	var temp []int

	for i, j := 0, 0; i < len(left) || j < len(right); {
		fmt.Println(i, j, len(left), len(right))

		if i == len(left) {

			temp = append(temp, right[j])
			j++
			continue
		}
		if j == len(right) {
			temp = append(temp, left[i])
			i++
			continue
		}

		if left[i] < right[j] {
			temp = append(temp, left[i])
			i++
		} else {
			temp = append(temp, right[j])
			j++
		}
	}

	return temp
}
