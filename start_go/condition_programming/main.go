package main

//if...else and switch

import "fmt"

func add(num1 int, num2 int) int {
	sum := num1 + num2

	return sum
}
func main() {

	a := 10
	b := 20

	sum := add(a, b)
	fmt.Println(sum)

	//conditional sign

	// >, >=, <, <=, ==
	//and => &&
	//or => ||
	//not => !

	// age := 5

	// if age > 18 {
	// 	fmt.Println("Your are eligible to be married")
	// } else if age < 18 {
	// 	fmt.Println("You are not eligible to be married, but you can love someone")
	// } else {
	// 	fmt.Println("you are just a teenager, not eligible to be married")
	// }

}
