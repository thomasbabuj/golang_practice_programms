//Checks whether an input alphabet is a vowel or not

package main

import "fmt"

var ch, status string

func main() {

	fmt.Println("Enter a character")
	_, err := fmt.Scanf("%s", &ch)

	if err != nil {
		fmt.Println("Error : ", err)
		return
	}

	if len(ch) != 1 {
		fmt.Println("Need to enter a character!")
		return
	}

	if ch == "a" || ch == "A" || ch == "e" || ch == "E" || ch == "i" ||
		ch == "I" || ch == "o" || ch == "O" || ch == "u" || ch == "U" {
		status = "a Vowel"
	} else {
		status = "Not an Vowel"
	}

	fmt.Println("The enter character '", ch, "' is ", status)
}
