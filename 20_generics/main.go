package main

import "fmt"

func printSlice[T int | string](items []T, name string) {
	fmt.Println(name)
	for _, item := range items {
		fmt.Println(item)
	}

}
func main() {
	nums := []int{1, 2, 3}
	num2 := []string{"a", "b"}
	printSlice(nums, "Aryan")
	printSlice(num2, "Gupta")
}
