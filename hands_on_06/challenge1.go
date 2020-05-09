//create a func with the identifier foo that returns an int
//create a func with the identifier bar that returns an int and a string
//call both funcs
//print out their results

package main

import "fmt"

func main() {
	fmt.Println(foo())
	x, y := bar()
	fmt.Println(x)
	fmt.Println(y)
}
func foo() int {
	return 42
}
func bar() (int, string) {
	a := 637
	b := "golang"
	return a, b
}
