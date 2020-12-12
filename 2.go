package main

import "fmt"

func main() {

	n := 4000000

	fib := []int{1, 1}

	sum := 0

	for i := 1; fib[i] < n; i++ {

		if fib[i]%2 == 0 {
			sum += fib[i]
		}

		fib = append(fib, fib[i]+fib[i-1])

	}

	fmt.Println(sum)

}
