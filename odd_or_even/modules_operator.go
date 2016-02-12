//Checking odd or even using modules operator

package main

import "fmt"

var input int64

func main() {
	fmt.Printf("Enter an integer: \n")
	fmt.Scanf("%d", &input)

	if input%2 == 0 {
		fmt.Printf("Even \n")
	} else {
		fmt.Printf("Odd \n")
	}
}
