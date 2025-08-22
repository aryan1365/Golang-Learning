package main

import "fmt"

type OrderStatus int

const (
	Recieved OrderStatus = iota
	Accepted
	Rejected
)

func changeStatus(status OrderStatus) {
	fmt.Println(status)
}

func main() {
	changeStatus(Rejected)
}
