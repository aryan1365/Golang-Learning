package main

import "fmt"

//by value
func changeNum(a int) {
	a = 4
	fmt.Println(a)
}
func changeByRef(a *int) {
	*a = 4
	fmt.Println(*a)
}
func main() {
	num := 2
	changeNum(num)
	fmt.Println("Without Reference", num)
	changeByRef(&num)
	fmt.Println("After Reference", num)
}
