package main

import "fmt"

func getPrimeFactorization (num int) map[int]int {
  factors := make(map[int]int)
  temp := num

  for i := 2; i <= temp; i++ {
    factorCtr := 0

    for temp%i == 0 {
      temp = temp/i
      factorCtr++
    }

    if factorCtr > 0 {
      factors[i] = factorCtr
    }
  }

  if (temp > 0) {
    factors[temp] = 1
  }

  return factors
}

func problem5 () {
  finalFactors := make(map[int]int)

  for i := 1; i <= 20; i++ {
    factors := getPrimeFactorization(i)

    for k, v := range factors {
      if v > finalFactors[k] {
        finalFactors[k] = v
      }
    }
  }

  answer := 1

  for k, v := range finalFactors {
    for i := 0; i < v; i++ {
      answer = answer * k
    }
  }
  
  fmt.Println(answer)
}
