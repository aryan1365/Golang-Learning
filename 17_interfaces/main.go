package main

import "fmt"

type paymenter interface {
	pay(amount float32)
}
type payment struct {
	gateway paymenter
}

func (p payment) makePay(amount float32) {
	p.gateway.pay(amount)
}

type razor struct{}

func (r razor) pay(amount float32) {
	fmt.Println("Paying using razorpay", amount)
}

type stripe struct{}

func (s stripe) pay(amount float32) {
	fmt.Println("Paying using stripe", amount)
}

func main() {
	rzpay := razor{}
	strp := stripe{}
	newPayment := []payment{
		{gateway: strp},
		{gateway: rzpay},
	}
	for _, p := range newPayment {
		p.makePay(100)
	}
}
