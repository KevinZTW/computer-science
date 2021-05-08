package main

import "fmt"

func Insertionsort(items []int) {
	var n = len(items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j = j - 1
		}
	}
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

func QuickSort(p []int) []int {
	if len(p) < 2 {
		return p
	}
	pivot := p[len(p)-1]
	start := -1
	for n, _ := range p {
		if p[n] < pivot {
			p[n]
		}
	}
}
