package main

import (
	"fmt"
	"time"
)

func main() {
	var a int = 10
	var b int = 20

	if a > b {
		fmt.Println("a is greater than b")
	} else if a < b {
		fmt.Println("a is less than b")
	} else {
		fmt.Println("a is equal to b")
	}



	if a > 5 && b > 15 {
		fmt.Println("Both conditions are true")
	}

	if a > 5 || b < 15 {
		fmt.Println("At least one condition is true")
	}

	if age := 13; age >= 18 {
		fmt.Println("You are an adult", age)
	} else if age >=12 && age < 18 {
		fmt.Println("You are a minor", age)
	} else {
		fmt.Println("You are a child", age)
	}

	//switch 
	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("Today is Monday")
	case "Tuesday":
		fmt.Println("Today is Tuesday")
	default:
		fmt.Println("Today is not Monday or Tuesday")
	}

	//multiple conditions switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a workday")
	}

	//type switch

	whoAmI := func(i interface{}) {
		switch v := i.(type) {
		case int:
			fmt.Printf("I am an int: %d\n", v)
		case string:
			fmt.Printf("I am a string: %s\n", v)
		default:
			fmt.Printf("I am of a different type: %T\n", v)
		}

	}

	whoAmI(42)
	whoAmI("Hello")
	whoAmI(3.14)
	// var x interface{} = 42
	// switch v := x.(type) {
	// case int:
	// 	fmt.Printf("x is an int: %d\n", v)
	// case string:
	// 	fmt.Printf("x is a string: %s\n", v)
	// default:
	// 	fmt.Printf("x is of a different type: %T\n", v)
	// }


	//GO DOSE NOT HAVE TERNARY OPERATOR BUT WE CAN USE IF ELSE STATEMENT INSTEAD 
	// (condition) ? expression1 : expression2
	// example
	// let text = (age < 18) ? "Minor" : "Adult";
	score := 85
	var grade string
	if score >= 90 {
		grade = "A"
	} else if score >= 80 {
		grade = "B"
	} else if score >= 70 {
		grade = "C"
	} else if score >= 60 {
		grade = "D"
	} else {
		grade = "F"
	}
	fmt.Printf("Your grade is %s\n", grade)

}
