//Using a function to check vowel

package main

import "fmt"

var input rune

func main() {
	fmt.Println("Enter a character")
	_, err := fmt.Scanf("%c", &input)

	if err != nil {
		fmt.Println("Error :", err)
	}

	if checkVowel(input) {
		fmt.Println("Vowel")
	} else {
		fmt.Println("Not an Vowel")
	}

}

func checkVowel(val rune) (status bool) {

	if val >= 'A' && val <= 'Z' {
		/*
		   Converting to lower case or use val = val + 32
		*/
		val = val + 'a' - 'A'
	}

	if val == 'a' || val == 'e' || val == 'i' || val == 'o' || val == 'u' {
		return true
	}

	return false

}
