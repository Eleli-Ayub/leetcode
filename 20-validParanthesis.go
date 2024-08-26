package main
func ValidParanthesis(s string)bool{
  //check if the length of the array is odd or even
  //loop over the array and check whether if each element and it's next value is value
  if len(s) == 0{
    return false 
  }

  if string(s[0]) == ")" || string(s[0]) == "}" || string(s[0]) == "]"{
    return false 
  }

  for i := 0; i <= len(s) -1 ; i++{
    if string(s[i]) == "("{
      return string(s[i+1]) == ")" 
    }else if string(s[i]) == "{"{
      return string(s[i+1]) == "}" 
    }else if string(s[i]) == "["{
      return string(s[i+1]) == "]" 
    }
  }

  return true
}
