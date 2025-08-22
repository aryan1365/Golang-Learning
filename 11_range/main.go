package main

import "fmt"

func main() {
	var nums = []int{1, 2, 3}
	var sum = 0
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
	}
	fmt.Println(sum)
	m := map[string]string{"name": "Aryan", "age": "21"}
	for k, v := range m {
		fmt.Println(k, v)
	}
	for i, c := range "golang" {
		fmt.Println(i, string(c))
	}
	for i, num := range nums {
		fmt.Println(i, num)
	}
}
