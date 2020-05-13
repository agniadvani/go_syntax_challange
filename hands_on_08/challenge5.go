//Fix the race condition you created in exercise #4 by using package atomic
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {

	var counter int64
	const gs = 100
	fmt.Println("Before waitgroup:", runtime.NumGoroutine())
	wg.Add(gs)
	for i := 0; i < 100; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("Counter: \t", atomic.LoadInt64(&counter))
			fmt.Println("waitgroup:", runtime.NumGoroutine())

			wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println("end waitgroup:", runtime.NumGoroutine())
	fmt.Println("end counter:", counter)
}
