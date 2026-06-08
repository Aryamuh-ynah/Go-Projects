package main

import (
	"fmt"
)

func main() {
	// for range
	slice := []string{"Go", "Python", "Java"}
	for index, value := range slice {
		fmt.Println("Index: %d, Value: %s\n", index, value)
	}

	// for range with array
	array := [3]int{1, 2, 3}
	for index, value := range array {
		fmt.Println("Index: %d, Value: %d\n", index, value)
	}

	// for range with map
	mapping := map[string]int{"Go": 1, "Python": 2, "Java": 3}
	for key, value := range mapping {
		fmt.Println("Key: %s, Value: %d\n", key, value)
	}

	// for range with string
	str := "Hello"
	for index, char := range str {
		fmt.Println(index, string(char))
	}

	// for range with channel
	// ch := make(chan int, 3)
	// ch <- 1
	// ch <- 2
	// ch <- 3
	// close(ch)

	// for value := range ch {
	// 	fmt.Printf("Value from channel: %d\n", value)
	// }


}
