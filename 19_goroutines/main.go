package main

import (
	"fmt"
	"sync"
)

// func task(id int) {
// 	fmt.Println("Doing task", id)
// }
// func main() {
// 	for i := 0; i <= 10; i++ {
// 		go task(i)
// 	}
// 	time.Sleep(time.Second * 1)
// }

// func f() {
// 	var i int
// 	for i = 0; i < 5; i++ {
// 		time.Sleep(10 * time.Millisecond)
// 		fmt.Print(i, " ")
// 	}
// }

// func main() {
// 	go f()
// 	f()
// }

// func PrintName(f string, l string, i int) {
// 	fmt.Println(f, l, i)
// }

// func main() {
// 	var i int
// 	go func() {
// 		for i = 0; i < 7; i++ {
// 			fmt.Print(i, " ")
// 			time.Sleep(100 * time.Millisecond)
// 		}
// 	}()
// 	time.Sleep(400 * time.Millisecond)
// 	PrintName("John", "Doe", i)
// 	time.Sleep(1 * time.Second)
// }

func task(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(id)
}
func main() {
	var wg sync.WaitGroup
	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go task(i, &wg)
	}
	wg.Wait()

}
