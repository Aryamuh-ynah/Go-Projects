package main

import (
	"fmt"
	"slices"
)

func main() {
	var nums [5]int

	for i := 0; i < len(nums); i++ {
		nums[i] = i*i*i*i
		
	}
	fmt.Println(nums)

	nums2 := [5]int{1, 8, 27, 64, 125}
	fmt.Println(nums2)


	var names []string 
	var n = make([]int, 3, 5)
	fmt.Println(len(names))
	fmt.Println(len(n))


	fmt.Println("capacity:", cap(n))
	fmt.Println("length:", len(n))

	names = append(names, "Humayra")
	names = append(names, "Fauzan")
	names = append(names, "Rizky")
	fmt.Println(names)

	n = append(n, 1)
	n = append(n, 2)
	n = append(n, 3)
	n = append(n, 4)
	fmt.Println(n)
	fmt.Println("capacity:", cap(n))
	fmt.Println("length:", len(n))



	//another way to create slice
	a := []int{}
	fmt.Println(a)
	fmt.Println(cap(a))
	fmt.Println(len(a))

	a = append(a, 1)
	a = append(a, 2)
	a = append(a, 3)
	fmt.Println(a)
	fmt.Println(cap(a))
	fmt.Println(len(a))

	// a[10]	= 11
	// fmt.Println(a)

	// copy slice
	b := make([]int, len(a))
	copy(b, a)
	fmt.Println(a, b)


	//slicing
	c := []int{1, 2, 3, 4, 5}
	fmt.Println("c:", c)

	d := c[1:4]
	fmt.Println("d(c[1:4])=", d)

	e := c[:3]
	fmt.Println("e(c[:3])=", e)

	f := c[2:5]
	fmt.Println("f(c[2:5])=", f)


	//2D slice
	g := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("g:", g)	


	x := []int{1, 2, 3, 4, 5}
	y :=[]int{6, 7, 8, 9, 10}

	fmt.Println("x:", x)
	fmt.Println("y:", y)

	fmt.Println("Is equal=", slices.Equal(x,y))

}