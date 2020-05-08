//Take the code from the previous exercise, then store the values of type person in a map with the key
//of last name. Access each value in the map. Print out the values, ranging over the slice.

package main

import "fmt"

type person struct {
	first    string
	last     string
	iceCream []string
}

func main() {
	p1 := person{
		first:    "Micheal",
		last:     "Scott",
		iceCream: []string{"chocoate", "banana"},
	}
	p2 := person{
		first:    "Jim",
		last:     "Halpert",
		iceCream: []string{"strawberry", "bubblegum"},
	}

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}
	fmt.Println(m)
	fmt.Println("----------")
	for k, v := range m {
		fmt.Println(k)
		fmt.Println("")
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, v := range v.iceCream {

			fmt.Println(i, v)
		}
		fmt.Println("")

	}
}
