package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("hello world")

	//fmt.Println(twoSumHash([]int{1, -1, 2, 4, 0}, 0))
	//fmt.Println(twoSumTwoPointer([]int{1, -1, 2, 4, 0}, 0))
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}, 0))
	// fmt.Println(threeSumClosest([]int{1, 0, -1, 2, 3, 3}, -2))
	//n := []int{3, 2, 1}
	//insertionsort(n)
	//fmt.Println(n)
}

func threeSum(nums []int, target int) [][]int{

	sort.Ints(nums)
	result, start, end, index, _, length := make([][]int, 0), 0, 0, 0, 0, len(nums)
	for index = 1; index < length; index++ {
       
        if index !=1 && nums[index] == nums[index-1]{
            continue
        }
		start, end = 0, length-1
		for start < index && end > index {
			fmt.Println(index, start,end)     
			if nums[start]+nums[end]+nums[index] > 0 {
				end--
				
			} else if nums[start]+nums[end]+nums[index] < 0 {
				start++
				
			} else {
				result = append(result, []int{nums[start], nums[end], nums[index]})
				break
			}
		}
	}
    return result
    
}

