//Fix the race condition you created in the previous exercise by using a mutex
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

func main() {

	counter := 0
	const gs = 100
	fmt.Println("Before waitgroup:", runtime.NumGoroutine())
	wg.Add(gs)
	for i := 0; i < 100; i++ {
		go func() {
			mu.Lock()
			v := counter
			v++
			counter = v
			fmt.Println("Counter\t\t", counter)
			fmt.Println("waitgroup:", runtime.NumGoroutine())
			mu.Unlock()
			wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println("end waitgroup:", runtime.NumGoroutine())
	fmt.Println("end counter:", counter)
}
