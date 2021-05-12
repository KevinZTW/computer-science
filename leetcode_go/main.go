package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")
}

func makeMatrix() {
	m, n := 3, 2

	//method 1 : use slice
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	//method 2 : use nest map
	dp2 := make(map[int]map[int]int, m)

	for i := 0; i < m; i++ {
		dp2[i] = make(map[int]int, n)
	}
}
