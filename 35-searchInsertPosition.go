package main
func SearchInsert(nums []int, target int) int {
	first, last := 0, len(nums)-1
	var index int
	//check if the target is the first or last element
	if target == nums[first] {
		return first
	} else if target == nums[last] {
		return last
	}
	if target > nums[len(nums)-1] {
		return len(nums)
	} else if target < nums[0] {
		return 0
	}
	//do a binary search on the array
  for i,v := range nums {
		mid := (first + last) / 2
		if target == nums[mid] {
			return mid
		}

		if target > v && target < nums[i+1] {
			return i + 1
		}
		if target > nums[mid] {
			first = mid
		} else {
			last = mid
		}
	}

	return index
}
