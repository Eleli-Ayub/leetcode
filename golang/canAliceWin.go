package main 

import (
  "sort"
)
func CanAliceWin(nums []int) bool{
  //sort the array
  sort.Ints(nums)
  //get the total of the numbers less than 9
  var singleDigitsSum, doubleDigitsSum int
  for _, v := range nums{
    if v <= 9{
      singleDigitsSum += v
    } else {
      doubleDigitsSum += v
    }
  }

  return singleDigitsSum != doubleDigitsSum
}
