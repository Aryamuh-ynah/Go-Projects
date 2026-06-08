package main

func add(a int, b int) int {
	return a + b
}

func multi()(string, string) {
	return "Hello", "Fauzan"
}	

// function in function
func outer(fn func(a int) string) {
	result := fn(10)
	println(result)
}

//returning a function
func returnFunc() func(a int) string {
	return func(a int) string {
		return "The value is:"
	}
}

func main(){
	result := add(5, 3)
	println("The result of addition is:", result)

	greeting, name := multi()
	println(greeting, name)

	outer(func(a int) string {
		return "The value passed to the function is: " + string(a)
	})


	returnedFn := returnFunc()
	println(returnedFn(10))	
}
