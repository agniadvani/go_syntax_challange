package main

import (
	"fmt"
	"sync"
)

func main() {
	eve := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	go foo(eve, odd)
	go bar(eve, odd, fanin)

	for i := range fanin {
		fmt.Println(i)
	}
	fmt.Println("About to exit")
}

func foo(eve, odd chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			eve <- i
		} else {
			odd <- i
		}
	}
	close(eve)
	close(odd)
}

func bar(eve, odd <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := range eve {
			fanin <- i
		}
	}()
	go func() {
		defer wg.Done()
		for i := range odd {
			fanin <- i
		}
	}()
	wg.Wait()
	close(fanin)
}
