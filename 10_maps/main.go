package main

import (
	"fmt"
	"maps"
)

func main() {
	//creating map
	m := map[string]string{}
	//setting elements
	m["name"] = "golang"
	fmt.Println(m["name"])

	m["name"] = "go"
	fmt.Println(m["name"])

	fmt.Println(len(m))
	m["type"] = "programming"

	fmt.Println(m)
	// delete(m, "type")
	// clear(m)
	// fmt.Println(m["name"])

	v, ok := m["name"]
	fmt.Println(v)
	if ok {
		fmt.Println("all Ok")
	} else {
		fmt.Println("Not ok")
	}
	map2 := map[string]string{"name": "go", "type": "programming"}
	fmt.Println(map2)
	fmt.Println(m)
	fmt.Println(maps.Equal(m, map2))
}
