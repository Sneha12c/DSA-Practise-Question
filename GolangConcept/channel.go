package main

import (
	"log"
	"sync"
)

func somefunc(num string, ch chan<- string, wg *sync.WaitGroup) {
	ch <- num
	log.Println(num)
	wg.Done()
}

func sliceToChannel(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, num := range nums {
			out <- num
		}
		close(out)
	}()
	return out
}

func square(ch <-chan int) <-chan int {
	final := make(chan int)
	go func() {
		for num := range ch {
			newnum := (num * num)
			final <- newnum
		}
		close(final)
	}()
	return final
}

func main() {

	// wg := &sync.WaitGroup{}
	// ch := make(chan string, 3)

	// wg.Add(3)
	// go somefunc("1", ch, wg)
	// go somefunc("2", ch, wg)
	// go somefunc("3", ch, wg)

	// go func() {
	// 	wg.Wait()
	// 	close(ch)
	// }()

	// for msg := range ch {
	// 	log.Println(msg)
	// }

	// mychannel := make(chan string)
	// myanotherchannel := make(chan string)

	// go func() {
	// 	mychannel <- "data"
	// }()

	// go func() {
	// 	myanotherchannel <- "cow"
	// }()

	// select {
	// case msgfrommychannel := <-mychannel:
	// 	log.Println(msgfrommychannel)
	// case msgfrommyanotherchannel := <-myanotherchannel:
	// 	log.Println(msgfrommyanotherchannel)
	// }

	// **** for select loop concurrency pattern

	ch := make(chan string, 3)
	chars := []string{"e", "t", "o"}

	for _, s := range chars {
		select {
		case ch <- s:
		}
	}

	close(ch)

	for res := range ch {
		log.Println(res)
	}

	log.Println("Hello")

	// input
	nums := []int{2, 34, 56, 78}
	// stage 1
	datatochannel := sliceToChannel(nums)
	// stage 2
	finalchannel := square(datatochannel)
	// stage 3
	for num := range finalchannel {
		log.Println(num)
	}

}

// Goroutine is a function which runs concurrently along with main goroutine
// It is a lightweight thread of execution managed entirely by goroutine
// Channels can be thought of as pipes using which Goroutines communicate.
// to make the communication asynchronous between goroutine, we need to use a buffered channel
// why buffered channel make communication synchronous because Unbuffered channel provide a gurantee
// that an exchange between goroutine send and receive data takes place instantley
// buffered channel follow a queue based approach , send and forget
// for select loop
// done channel
// pipeline

// Concurrency Part 2
// Generators -
