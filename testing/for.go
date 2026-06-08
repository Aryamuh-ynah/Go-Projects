package main

import "fmt"

func main() {
	println("Hello World")

	//while loop using for
	i := 0
	for i < 3 {
		fmt.Println(i)
		i++
	}

	//infinite loop
	// for {
	// 	fmt.Println("This is an infinite loop")
	// }

	// for loop
	for i := 0; i < 4; i++ {
		fmt.Println(i)
	}

	// for loop with break
	for i := 0; i < 7; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

	// for loop with continue
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	// for loop with multiple variables
	for i, j := 0, 10; i < 5; i, j = i+1, j-1 {
		fmt.Printf("i: %d, j: %d\n", i, j)
	}

	// for range
	for i:= range 4{
		fmt.Println(i)
	}


}
