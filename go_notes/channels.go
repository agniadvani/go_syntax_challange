package main

import "fmt"

func main() {
	ch := make(chan int)
	c := make(chan int)
	go func() {
		ch <- 12 + 14
	}()
	go func() {
		c <- <-ch + 21

	}()
	f := <-c + 12
	fmt.Println(f)
}
