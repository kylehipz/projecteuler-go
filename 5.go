package main

import "fmt"

func check(num int) bool {
	for i := 2; i <= 20; i++ {
		if num%i != 0 {
			return false
		}
	}

	return true
}

func main() {

	found := false

	num := 100

	for !found {

		if check(num) {
			fmt.Println(num)
			break
		}

		num++
	}

}
