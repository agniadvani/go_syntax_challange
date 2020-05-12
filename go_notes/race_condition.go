package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	counter := 0
	const gs = 100
	var wg sync.WaitGroup
	fmt.Println("Cpu:", runtime.NumCPU())

	wg.Add(gs)
	for i := 0; i < 100; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("routines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("routines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
