package main

import (
	"fmt"
	"time"
)

type Order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
}

func NewOrder(id string, amount float32, status string) *Order {
	myOrder := Order{
		id:     id,
		amount: amount,
		status: status,
	}
	return &myOrder
}
func (O *Order) changeStatus(status string) {
	O.status = status
}

func main() {
	order1 := NewOrder("1", 30.34, "Recieved")
	order2 := NewOrder("2", 100.00, "Paid")

	order1.createdAt = time.Now()
	order1.changeStatus("Delievered")
	fmt.Println(order1)
	fmt.Println(order2)

	language := struct {
		name   string
		isGood bool
	}{"go", true}
	fmt.Println(language.isGood)
}
