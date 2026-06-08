package main

// Variadic function (accepts n number of arguments)
func variadicSum(nums ...int) int {
	sum := 0
	for _, num := range nums {

		sum += num
	}
	return sum
}

func main() {
	result := variadicSum(1, 2, 3, 4, 5)
	println("The sum of the numbers is: ", result)

	// variadic function with no arguments
	result2 := variadicSum()
	println("The sum of no numbers is", result2)

	// variadic function with one argument
	result3 := variadicSum(10)
	println("The sum of one number is:", result3)

	nums := []int{1, 2, 3, 4, 5}
	result4 := variadicSum(nums...)
	println("The sum of the numbers in the slice is:", result4)
}