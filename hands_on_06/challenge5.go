//Create a func which returns a func
//assign the returned func to a variable
//call the returned func

package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := addOdd(a...)
	fmt.Println(b)

}

func add(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func addOdd(y ...int) int {
	xi := []int{}
	for _, v := range y {
		if v%2 != 0 {
			xi = append(xi, v)
		}
	}
	total := add(xi...)
	return total
}

