/*
Program ask user to input an integer and then
prints it.
*/
package main

import "fmt"

var userInput int

func main() {
	fmt.Println("Enter an Integer")
	_, err := fmt.Scanf("%d", &userInput)

	if err != nil {
		fmt.Println("Error ", err)
		return
	}

	fmt.Printf("Integer that you have entered is %d \n", userInput)
}
