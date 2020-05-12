//Using goroutines, create an incrementer program
//have a variable to hold the incrementer value
//launch a bunch of goroutines
//each goroutine should
//read the incrementer value
//store it in a new variable
//yield the processor with runtime.Gosched()
//increment the new variable
//write the value in the new variable back to the incrementer variable
//use waitgroups to wait for all of your goroutines to finish
//the above will create a race condition.
//Prove that it is a race condition by using the -race flag

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	counter := 0
	const gs = 100
	fmt.Println("Before waitgroup:", runtime.NumGoroutine())
	wg.Add(gs)

	for i := 0; i < 100; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("waitgroup:", runtime.NumGoroutine())
		fmt.Println("counter\t:", counter)
	}
	wg.Wait()
	fmt.Println("end waitgroup:", runtime.NumGoroutine())
	fmt.Println("end counter:", counter)
}
