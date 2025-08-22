package main

import "fmt"

func add(a int, b int) int {
	return a + b
}
func sub(a int, b int) int {
	return a - b
}

// func getLan() (string, string, string) {
// 	return "golang", "js", "c"
// }
var getLan = func() (string, string, string) {
	return "golang", "js", "react"
}

func processIt(fn func(a int) int) int {
	return fn(2)
}
func Pr() func(a int) int {
	return func(a int) int {
		return 3
	}
}
func main() {
	fmt.Println(add(3, 3))
	fmt.Println(sub(4, 3))
	lang1, lang2, lang3 := getLan()
	// fmt.Println(lang1, lang2, lang3)
	fmt.Println(lang1, lang2, lang3)
	var fn = func(a int) int {
		return 3
	}
	fmt.Println(processIt(fn))
}
