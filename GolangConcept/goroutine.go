package main

import (
	"fmt"
	"sync"
)

func task(i int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println("Task done", i)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go task(i, &wg)
	}
	// time.Sleep(time.Second * 2)
	wg.Wait()
}
