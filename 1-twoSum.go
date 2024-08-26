package main

//Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
//You may assume that each input would have exactly one solution, and you may not use the same element twice.
//You can return the answer in any order.

func TwoSum(nums []int, target int) []int {
	var pairsMap map[int]int

	for i, v := range nums {
		difference := target - v

		if index, found := pairsMap[difference]; found {
			return []int{index, i}
		}

		pairsMap[v] = i
	}

	return nil
}
