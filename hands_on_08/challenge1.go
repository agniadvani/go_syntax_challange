//in addition to the main goroutine, launch two additional goroutines
//each additional goroutine should print something out
//use waitgroups to make sure each goroutine finishes before your program exists.

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	fmt.Println("Before foo: ", runtime.NumGoroutine())
	wg.Add(1)
	go foo()
	fmt.Println("After foo: ", runtime.NumGoroutine())
	wg.Add(1)
	go bar()
	fmt.Println("After bar: ", runtime.NumGoroutine())
	wg.Wait()
}

func foo() {
	fmt.Println("This is foo.")
	wg.Done()
}

func bar() {
	fmt.Println("This is bar.")
	wg.Done()
}
