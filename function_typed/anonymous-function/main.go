package main

// import "fmt"

func missingNumber(a []int, n int) {
	x1 := 0
	x2 := 0

	for i := 0; i < n-1; i++ {
		x1 = x1 ^ a[i]
	}
	for i := 1; i <= n; i++ {
		x2 = x2 ^ i
	}

	println(x1 ^ x2)
}

func main() {
	//anonymous function
	// func(a int, b int) {
	// 	c := a + b
	// 	fmt.Println(c)
	// }(5, 7)

	arr1:= [4]int{1,3,5,4}
	missingNumber(arr1[:], 5)
}
