package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var counter = 0
	const gs = 100
	var wg sync.WaitGroup

	fmt.Println("routines:", runtime.NumGoroutine())
	wg.Add(gs)
	var mu sync.Mutex
	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("routines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("routines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)

}
