package main

import (
	"fmt"
)

func FourSum(nums []int, target int) [][]int {
	// create map of nums
	// loop over array
	// for each value get the difference with the target
	// check if a value less than or equal to the remainder exists
	// append each value that adds up to the target to a given array
	// append the new array to result
	numsMap := make(map[int]int, len(nums))
	for i, v := range nums {
		numsMap[i] = v
	}

	for i, v := range nums {
		rem := target - v
		fmt.Println(v)
		for j, jv := range nums {
			if i == j {
				continue
			}
			if jv <= rem {
			}
		}
	}
	return [][]int{}
}
