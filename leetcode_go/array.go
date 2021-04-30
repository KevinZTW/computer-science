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
       
        if index !=0 && nums[index] == nums[index-1]{
            
            continue
        }
		start, end = 0, length-1
		for start < index && end > index {
			

            
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

//1. Use sort for two pointer trick


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

//Sort and use two pointer start from the head and the end

func threeSumClosest(nums []int, target int) int {
    sort.Ints(nums)
    index, start, end, length,diff, result := 0, 0, 0, len(nums), 0,0
    fmt.Println(length)
    
    for index = 1; index < length; index++{
		start, end = 0, length-1
		fmt.Println(index, start, end, diff)
		
		
		for start< index && end >index{
			fmt.Println(start, index, end,nums[start], nums[index], nums[end])
			if nums[start]+nums[index]+nums[end]< target{
				currentDiff := nums[start]+nums[index]+nums[end]-target
				if currentDiff <0 {
					currentDiff = 0-currentDiff
				}
				if diff ==0 || currentDiff < diff { 
					diff = currentDiff
					result = nums[start]+nums[index]+nums[end]
				}

				start ++	
			}else if nums[start]+nums[index]+nums[end] > target {
				currentDiff := nums[start]+nums[index]+nums[end]-target
				if currentDiff <0 {
					currentDiff = 0-currentDiff
				}
				if diff ==0 || currentDiff < diff { 
					diff = currentDiff
					result = nums[start]+nums[index]+nums[end]
				}
				end --
			}else{
				return target
			}
		}
	} 
	
	return result
}