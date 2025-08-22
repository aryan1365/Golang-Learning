package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for i := range nums {
		total += nums[i]
	}
	return total
}
func main() {
	total := sum(23, 33344, 5, 56, 7, 8)
	num := []int{1, 2, 3, 4}
	fmt.Println(total)
	res := sum(num...)
	fmt.Println(res)
}
