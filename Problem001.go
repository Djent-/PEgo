//Problem 1
//Sum the list of all natural numbers below 1000 that are
//multiples of 3 or 5

package main

import "fmt"

func main() {
	count := 0
	for x := 0; x < 1000; x++ {
		if x % 5 == 0 || x % 3 == 0 {
			count = count + x
		}
	}
	fmt.Println(count)
}