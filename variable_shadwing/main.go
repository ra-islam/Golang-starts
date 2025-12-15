package main

// important interview question

import "fmt"

//global scope
var a = 10

func main() {
	//local scope
	age := 30

	if age > 18 {
		//block scope
		var a = 47 // shadowing

		fmt.Println(a)

		//first find the a in block scope if find then print and shadowing the value of global scope. if not found then check in the local scope
	}

	fmt.Println(a)
	// check in the local scope if not find then check in the global scope
}
