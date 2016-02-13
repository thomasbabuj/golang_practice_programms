//Add digits of a number.
package main

import "fmt"

var input int

func main() {
	fmt.Println("Enter an integer")
	fmt.Scanf("%d", &input)

	var remainder, sum int

	t := input

	for t != 0 {
		remainder = t % 10
		sum += remainder
		t = t / 10
	}

	fmt.Printf("Sum of a given digit (%d) = %d \n", input, sum)
}
