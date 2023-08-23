package main

import "fmt"

func isPalindrome (num int) bool {
  reversed := 0
  temp := num

  for temp > 0 {
    lastDigit := temp%10
    reversed = reversed * 10 + lastDigit

    temp = temp/10
  }

  return num == reversed
}

func problem4 () {
  largest := 0

  for i := 100; i <= 999; i++ {
    for j := 100; j <= 999; j++ {
      product := i*j
      if isPalindrome(product) && product > largest {
        largest = product
      }
    }
  }

  fmt.Println(largest)
}
