
package main

import "fmt"

func problem10 () {
  sum := 0

  for i := 2; i < 2*1e6; i++ {
    if isPrime(i) {
      sum += i
    }
  }

  fmt.Println(sum)
}
