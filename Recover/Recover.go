package main

import "fmt"

func process() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()
	fmt.Println("Start Process")
	panic("Something Went Wrong")
}

func main() {
	process()
	fmt.Println("Process Completed")
}
