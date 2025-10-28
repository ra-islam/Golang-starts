package main

import "fmt"

func printWelcomeMessage() {
	//Print Welcome message
	fmt.Println("Welcome to the Mini Project in Go!")
}

func getUserName() string {
	// Get user input
	var name string
	fmt.Println("Enter your name:")
	fmt.Scanln(&name)

	return name
}

func getTwoNumbers() (int, int) {
	var num1 int
	var num2 int
	fmt.Println("Enter first number:")
	fmt.Scanln(&num1)
	fmt.Println("Enter second number:")
	fmt.Scanln(&num2)

	return num1, num2
}

func add(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}

func displayResult(name string, sum int) {
	fmt.Printf("Hello, %s! The sum is %d.\n", name, sum)
}

func goodByeMessage() {
	fmt.Println("Thank you for using the Mini Project in Go. Goodbye!")
}

func main() {
	printWelcomeMessage()
	name := getUserName()
	num1, num2 := getTwoNumbers()
	// Perform addition
	sum := add(num1, num2)
	displayResult(name, sum)
	goodByeMessage()
}
