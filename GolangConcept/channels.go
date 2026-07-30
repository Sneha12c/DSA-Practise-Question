package main

import (
	"fmt"
	"time"
)

func processNum(numChan chan int) {
	for num := range numChan {
		fmt.Println("processing function", num)
		time.Sleep(time.Second)
	}
}

func sum(result chan int, a int, b int) {
	sumres := a + b
	result <- sumres
}

// Go routine Synchronization using channels
// func task(done chan bool) {
// 	defer func() { done <- true }()

// 	fmt.Println("Procesing ...")
// }

func emailSender(emailChan chan string, done chan bool) {
	defer func() { done <- true }()
	for email := range emailChan {
		fmt.Println("sending email to", email)
	}
}

func main() {

	chan1 := make(chan string)
	chan2 := make(chan int)

	go func() {
		chan1 <- "ping"
	}()

	go func() {
		chan2 <- 45
	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("Chan1 is processed", chan1Val)
		case chan2Val := <-chan2:
			fmt.Println("Chan2 is processed", chan2Val)
		}
	}

	// emailchan := make(chan string, 10)
	// done := make(chan bool)

	// go emailSender(emailchan, done)

	// for i := 0; i < 10; i++ {
	// 	emailchan <- fmt.Sprintf("%d@gmail.com", i)
	// }

	// close(emailchan) // buffered channel should be closed
	// <-done

	// Unbuffered channels - receiving and sending are blocking code
	// done := make(chan bool)

	// go task(done)
	// <-done

	// result := make(chan int)

	// go sum(result, 4, 5)

	// res := <-result

	// fmt.Println(res)

	// numChan := make(chan int)

	// go processNum(numChan)

	// for {
	// 	numChan <- rand.Intn(100)
	// }

	// messagechannel := make(chan string)

	// messagechannel <- "ping"

	// msg := <-messagechannel

	// fmt.Println(msg)

}
