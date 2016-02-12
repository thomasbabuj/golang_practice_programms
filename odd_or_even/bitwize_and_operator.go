//Check odd or even using bitwise operator
//an example consider binary of 7 (0111) when
//we perform 7 & 1 the result will be one and
//you may observe that the least significant
//bit of every odd number is 1,
//so ( odd_number & 1 ) will be one always
//and also ( even_number & 1 ) is zero.

package main

import "fmt"

var input int

func main() {
	fmt.Printf("Enter an Integer \n")
	fmt.Scanf("%d", &input)

	if input&1 == 1 {
		fmt.Printf("Odd \n")
	} else {
		fmt.Printf("Even \n")
	}

}
