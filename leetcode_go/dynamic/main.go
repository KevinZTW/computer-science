//重疊子問題：在窮舉的過程中（比如通過遞迴），存在重複計算的現象；
//無後效性：子問題之間的依賴是單向性的，某階段狀態一旦確定，就不受後續決策的影響；
//最優子結構：子問題之間必須相互獨立，或者說後續的計算可以通過前面的狀態推匯出來。

// 53. Maxium Subarray
package main

import "fmt"

func main() {
	fmt.Println(findMaxSubarray([]int{2, -1, -2, 3, -1, 4, 5}))
	fmt.Println(coinChange([]int{1, 2, 4}, 11))
}

func findMaxSubarray(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	res := nums[0]

	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		if res < dp[i] {
			res = dp[i]
		}
	}
	return res
}

//[]int{1, 2, 5}, 11, 3   5+5+1
func coinChange(coins []int, amount int) int {

	dp := make(map[int]int, amount+1)

	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}

	for i := 0; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				if dp[i] > dp[i-coins[j]]+1 {
					dp[i] = dp[i-coins[j]] + 1
				}
			}
		}
	}
	fmt.Println(dp)

	return dp[amount]
}
