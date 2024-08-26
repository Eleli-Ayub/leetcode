package main

import (
  "fmt"
  "sort"
)
func LongestCommonPrefix(strs [] string) string{
  prefix := ""
  sort.Strings(strs)
  for i, v := range strs[0]{
    if string(v) == string(strs[len(strs)-1][i]){
      prefix = strs[0][:i+1]
    }else{
      break
    }
  }
  fmt.Println(prefix)
  return prefix
}
