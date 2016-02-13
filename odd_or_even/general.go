//Check odd or even without using bitwise or modules operator
package main

import "fmt"

var input int

func main() {
	fmt.Println("Enter an Integer")
	fmt.Scanf("%d", &input)

	if (input/2)*2 == input {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
