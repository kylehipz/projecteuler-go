package main

import (
	"fmt"
	"math"
)

func checkPrime(num int) bool {

	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func main() {

	n := 600851475143

	largest := 0.0

	for i := 2; float64(i) < math.Sqrt(float64(n)); i++ {
		if n%i == 0 && checkPrime(i) {
			largest = math.Max(float64(largest), float64(i))
		}
	}
	fmt.Println(largest)
}
