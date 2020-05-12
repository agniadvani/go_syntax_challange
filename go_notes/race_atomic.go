package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64
	const gs = 100
	var wg sync.WaitGroup

	fmt.Println("routines:", runtime.NumGoroutine())
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("Counter:\t", atomic.LoadInt64(&counter))
			wg.Done()
		}()
		fmt.Println("routines:", runtime.NumGoroutine())

	}
	wg.Wait()

	fmt.Println("count:", counter)

}
