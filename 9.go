package main

import "fmt"

func isPythagoreanTriplet (a int, b int, c int) bool {
  return (a*a) + (b*b) == (c*c)
}

func problem9 () {
  for a := 1; a < 1000; a++ {
    for b := 1; b < 1000; b++ {
      for c := 1; c < 1000; c++ {
        if isPythagoreanTriplet(a, b, c) && a+b+c == 1000 {
          fmt.Println(a, b, c)
          fmt.Println(a*b*c)
        }
      }
    }
  }
}
