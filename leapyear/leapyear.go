//Program to check leap year
package main

import "fmt"

var year int

func main() {
	fmt.Printf("Enter a year to check if it is a leap year \n")
	_, err := fmt.Scanf("%d", &year)

	if err != nil {
		fmt.Println("Error : ", err)
	}

	if year%400 == 0 {
		fmt.Printf("%d is a leap year.\n", year)
	} else if year%100 == 0 {
		fmt.Printf("%d is not leap year.\n", year)
	} else if year%4 == 0 {
		fmt.Printf("%d is a leap year. \n", year)
	} else {
		fmt.Printf("%d is not a leap year. \n", year)
	}
}
