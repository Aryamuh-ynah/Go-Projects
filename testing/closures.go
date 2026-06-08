package main

import "fmt"


func counter() func() int {

	var count int =0 

	return func() int {
		count++
		return count
	}
}

func main() {

	inc := counter()

	fmt.Println(inc()) // Output: 1
	fmt.Println(inc()) // Output: 2
	fmt.Println(inc()) // Output: 3

	// Create another counter
	anotherInc := counter()
	fmt.Println(anotherInc()) // Output: 1
	fmt.Println(anotherInc()) // Output: 2
	
}