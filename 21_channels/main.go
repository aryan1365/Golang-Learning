package main

import "fmt"

// sending
// func processNum(numChan chan int) {
// 	for num := range numChan {
// 		fmt.Println("Proccesing number", num)
// 		time.Sleep(time.Second)
// 	}
// }
// func main() {
// 	//channels are blocking
// 	// messageChan := make(chan string)
// 	// messageChan <- "ping"
// 	// message := <-messageChan
// 	// fmt.Println(message)

// 	numChan := make(chan int)
// 	go processNum(numChan)
// 	for {
// 		numChan <- rand.Intn(200)
// 	}
// 	// time.Sleep(time.Second * 2)
// }

// func sum(result chan int, num1 int, num2 int) {
// 	sum := num1 + num2
// 	result <- sum
// }
// func main() {
// 	result := make(chan int)
// 	go sum(result, 5, 5)
// 	res := <-result
// 	fmt.Println(res)
// }

// func task(done chan bool) {
// 	defer func() { done <- true }()
// 	fmt.Println("Processing")

// }
// func main() {
// 	done := make(chan bool)
// 	go task(done)
// 	<-done
// }

// buffered chaneel
// func sendEmail(email chan string, done chan bool) {
// 	defer func() { done <- true }()
// 	for em := range email {
// 		fmt.Println(em)
// 		time.Sleep(time.Millisecond * 5)
// 	}
// }
// func main() {
// 	emailChan := make(chan string, 100)
// 	done := make(chan bool)
// 	// emailChan <- "1@gmail"
// 	// emailChan <- "2@gmail"
// 	// fmt.Println(<-emailChan)
// 	// fmt.Println(<-emailChan)
// 	go sendEmail(emailChan, done)
// 	for i := 0; i < 100; i++ {
// 		emailChan <- fmt.Sprintf("%d@gmail.com", i)
// 	}
// 	close(emailChan)
// 	<-done
// }

func main() {
	chan1 := make(chan int)
	chan2 := make(chan string)
	go func() {
		chan1 <- 10
	}()
	go func() {
		chan2 <- "aryan"
	}()
	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("recieved on chan1", chan1Val)
		case chan2Val := <-chan2:
			fmt.Println("recieved on chan2", chan2Val)
		}
	}
}
