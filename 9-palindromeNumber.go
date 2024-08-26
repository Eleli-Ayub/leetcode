package main
import "strconv"
func PalindromeNumber(x int) bool{
  numberAsString := strconv.Itoa(x)
  index := 0
  lastIndex := len(numberAsString) - 1
  for index <= lastIndex{
    if numberAsString[index] != numberAsString[lastIndex]{
      return false
      break
    }
    index ++
    lastIndex --
  }
  return true
}
