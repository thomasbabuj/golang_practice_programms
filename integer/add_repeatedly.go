//Add two numbers repeatedly

package main

import "fmt"

var one, two int
var status string

func main() {
	for {
		fmt.Println("Input two integers")
		fmt.Scanf("%d %d", &one, &two)

		result := one + two
		fmt.Printf("(%d) + (%d) = (%d) \n", one, two, result)

		fmt.Println("Do you wish to add more numbers (y/n)")
		_, err := fmt.Scanf("%s", &status)

		if err != nil {
			fmt.Println("Error : ", err)
		}

		if status == "y" || status == "Y" {
			continue
		} else {
			fmt.Println("Thanks for using our system!!")
			break
		}

	}
}
