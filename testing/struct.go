package main

import (
	"fmt"
	"time"
)

type order struct {
	id string
	amount float32
	status string
	createdAt time.Time


}


func main(){

	order1 := order{
		id:"1",
		amount:399,
		status: "Received", 
	}

	order1.createdAt = time.Now()

	fmt.Println(order1.amount)
	fmt.Println("Order struct ", order1)

	order2 := order{
		id: "2",
		amount: 122,
		status: "Delivered",
		createdAt: time.Now(),

	}

	fmt.Println("Order 2", order2)
}