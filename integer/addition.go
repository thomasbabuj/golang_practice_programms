//Add two numbers.
package main

import "fmt"

var a, b int

func main() {

	fmt.Println("Enter two numbers to add")
	_, err := fmt.Scanf("%d %d", &a, &b)

	if err != nil {
		fmt.Println("Error : ", err)
	}

	c := a + b
	fmt.Println("Sum of entered numbers = ", c)
}
