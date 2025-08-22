package main

import "fmt"

func main() {

	// age := 15
	// if age >= 18 {
	// 	fmt.Println("You are an adult")
	// } else if age >= 12 {
	// 	fmt.Println("Teen")
	// } else {
	// 	fmt.Println("Not an adult")
	// }

	var role = "Admin"
	var hasPerm = false
	if role == "Admin" && hasPerm {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
	if age := 15; age >= 12 && age < 18 {
		fmt.Println("Teenage")
	}

}
