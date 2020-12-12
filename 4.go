package main

import (
	"fmt"
)

func isPalindrome(num int) bool {
	tmp := num
	s := 0

	for num > 0 {
		s = (s * 10) + (num % 10)

		num /= 10
	}

	if s == tmp {
		return true
	}

	return false
}

func main() {

	largest := 0

	for i := 100; i <= 999; i++ {
		for j := 100; j <= 999; j++ {
			prod := i * j

			if isPalindrome(prod) && prod > largest {
				largest = prod
			}
		}
	}

	fmt.Println(largest)
}
