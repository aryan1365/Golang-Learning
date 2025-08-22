package main

import (
	"fmt"
	"slices"
)

func main() {
	//unitialized slice is nill
	var nums []int
	fmt.Println(len(nums))

	var arr = make([]int, 2, 5)
	//capacity->maximum no of element that can fit
	fmt.Println(len(arr))
	nums = append(nums, 3)
	fmt.Println(nums)
	arr[0] = 2
	arr = append(arr, 1, 2, 3, 4, 5)
	fmt.Println(arr)

	// temp := []int{}

	//copyfn
	nums = append(nums, 4)
	var nums2 = make([]int, len(nums))
	copy(nums2, nums)
	fmt.Println(nums2)
	sliceNum := []int{1, 2, 3}
	fmt.Println(sliceNum[0:2])
	fmt.Println(slices.Equal(nums, nums2))
}
