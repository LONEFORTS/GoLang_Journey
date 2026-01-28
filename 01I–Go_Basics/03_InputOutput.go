package main

import "fmt"

func main() {
	// Declare variables
	var name string
	var age int

	// Taking input from user
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)

	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)

	// Displaying output
	fmt.Println("\n--- User Details ---")
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
}