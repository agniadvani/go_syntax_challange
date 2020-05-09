// addition and multiplication of even numbers.

package main

import "fmt"

func main() {
	fmt.Println(evenOps(sum, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}...))
	fmt.Println(evenOps(multi, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}...))

}

func sum(x ...int) int {
	n := 0
	for _, v := range x {
		n += v
	}
	return n
}
func multi(x ...int) int {
	n := 1
	for _, v := range x {
		n *= v
	}
	return n
}

func evenOps(s func(x ...int) int, y ...int) int {
	var xi []int
	for _, v := range y {
		if v%2 == 0 {
			xi = append(xi, v)
		}
	}
	total := s(xi...)
	return total
}
