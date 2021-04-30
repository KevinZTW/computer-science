package main

import (
	"fmt"
	"sort"
)

func twoSumHash(nums []int, target int) []int {
	numsMap := make(map[int]int)

	for _, n := range nums {
		t := target - n
		_, ok := numsMap[t]
		if ok {
			return []int{n, t}
		} else {
			numsMap[n] = n
		}
	}
	return nil
}

func twoSumTwoPointer(nums []int, target int) []int {

	return nil
}

func threeSum(nums []int, target int) {
	sort.Ints(nums)
	result, start, end, index, _, length := make([][]int, 0), 0, 0, 0, 0, len(nums)
	for index = 1; index < length; index++ {
		start, end = 0, length-1
		for start < index && end > index {
			fmt.Println(index, start, end)
			if nums[start]+nums[end]+nums[index] > target {
				end--
				continue
			} else if nums[start]+nums[end]+nums[index] < target {
				start++
				continue
			} else {
				result = append(result, []int{nums[start], nums[end], nums[index]})
				fmt.Println(result)
				break
			}
		}
	}
}
