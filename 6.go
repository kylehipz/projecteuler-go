package main

import "fmt"

func problem6 () {
  sumOfSquares := 0
  squareOfSum := 0

  for i := 1; i <= 100; i++ {
    sumOfSquares += i*i
    squareOfSum += i
  }

  squareOfSum = squareOfSum * squareOfSum
  difference := squareOfSum - sumOfSquares

  fmt.Println(difference)
}
