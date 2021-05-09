package main

import "fmt"

func main() {
	nums := []int{11, 2, -4, 3, -1, 6}
	fmt.Println(Enumeration(nums))
	fmt.Println(Dynamic(nums))
}

// for array A1, A2.....An find 1<=i<=j<=n that Ai + Ai+1 ...Aj has the biggest sum

//1. Enumeration
func Enumeration(nums []int) int {
	l := len(nums)
	best := nums[0]
	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			sum := 0
			for k := i; k <= j; k++ {
				sum += nums[k]
			}
			if sum > best {
				best = sum
			}
		}

	}

	return best
}

//2. Dynamic Programming
func Dynamic(nums []int) int {
	dp := make(map[int]int, len(nums))
	dp[0] = nums[0]
	l := len(nums)
	max := nums[0]
	for i := 1; i < l; i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
	}
	for j := 0; j < l; j++ {
		if dp[j] > max {
			max = dp[j]
		}
	}
	return max
}

//3. Greedy

//4. Devide and conquer
