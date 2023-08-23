package main

import "fmt"

func problem7 () {
  ctr := 1
  cur := 2

  for ctr < 10001 {
    cur++
    if isPrime(cur) {
      ctr++
    }
  }

  fmt.Println(cur)
}
