//Check vowel using switch statement
package main

import "fmt"

var input string

func main() {
	fmt.Println("Enter a character")
	fmt.Scanf("%s", &input)

	switch input {
	case "a", "A", "e", "E", "i", "I", "o", "O", "u", "U":
		fmt.Println("The given character is an Vowel")
	default:
		fmt.Println("The given character is not an Vowel!!")
	}
}
