package main

import "fmt"

func isPrime (num int) bool {
  for i := 2; i*i <= num; i++ {
    if num%i == 0 {
      return false
    }
  }

  return true
}

func problem3 () {
  num := 600851475143
  largest := 0

  for i := 2; i*i <= num; i++ {
    if isPrime(i) && num%i == 0 {
      largest = i
    }
  }

  fmt.Println(largest)
}
