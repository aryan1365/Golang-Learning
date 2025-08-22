package main

import "fmt"

func main() {
	var nums [5]int
	nums[0] = 1
	fmt.Println(len(nums))
	fmt.Println(nums[0])

	arr := [3]int{1, 2, 3}
	fmt.Println(arr)

	temp := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(temp)
}
