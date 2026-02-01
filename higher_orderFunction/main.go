package main

import "fmt"

func processOperation(a int, b int, op func(p int, q int)) {
	op(a, b)
}

func add(x int, y int) {
	z := x + y
	fmt.Println(z)
}
func main() {
	processOperation(2, 5, add)
}

/* 
Higher Order function -
1. Accept Function as parameter -- callback function
2. return a function
3. 1 and 2 both 
*/
