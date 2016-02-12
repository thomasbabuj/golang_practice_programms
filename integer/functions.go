//Adding numbers using functions
package main

import "fmt"

var first, second, sum int64

func main() {
	fmt.Printf("Enter two numbers to add :\n")
	fmt.Scanf("%d %d", &first, &second)
	sum = addition(first, second)
	fmt.Printf("Your total : %d \n", sum)
}

func addition(first, second int64) int64 {
	return first + second
}
