package main

import "fmt"

func main() {
	i := 4
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Three")
	default:
		fmt.Println("Four")
	}

	//type switch
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("Integer")
		case string:
			fmt.Println("String")
		case bool:
			fmt.Println("Boolean")
		default:
			fmt.Println("other", t)
		}
	}
	whoAmI(true)
}
