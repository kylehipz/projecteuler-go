package main

import "fmt"

func problem2 () {
  sum := 0

  prev := 1
  cur := 2

  for cur < 4*1e6 {
    if cur%2 == 0 {
      sum += cur
    }

    next := prev + cur

    prev = cur
    cur = next
  }

  fmt.Println(sum)
}
