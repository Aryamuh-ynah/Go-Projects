package main

import (
	"fmt"
	"maps"
)

func main() {

	m := make(map[string]string)

	m["a"] = "Cat"
	m["b"] = "Dog"
	m["c"] = "Fish"

	//get

	fmt.Println(m["a"], m["b"], m["k"])

	n := map[string]int{
		"one": 1,
		"two": 2,
		"three": 3,
	}

	fmt.Println(n["one"], n["two"], n["ten"])
	fmt.Println(len(m))

	delete(m, "b")
	fmt.Println(m)

	//check if key exists
	value, exists := m["a"]
	if exists {
		fmt.Println("Key exists:", value)
	} else {
		fmt.Println("Key does not exist")
	}

	k := maps.Clone(m)

	fmt.Println(maps.Equal(m, k))


	clear(n)
	fmt.Println(n)

}