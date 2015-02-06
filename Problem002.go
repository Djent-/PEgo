//Problem 2
//Find the sum of all all even Fibonacci numbers less than 4,000,000

package main

import "fmt"

func main() {
	a := 0
	b := 1
	sum := 0
	for {
		if b % 2 == 0 {
			sum += b
		}
		a, b = b, b + a
		if b > 4000000 {
			break
		}
	}
	fmt.Println(sum)
}