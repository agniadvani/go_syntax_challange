//Create your own type “person” which will have an underlying type of “struct” so that it can store
//the following data:

//first name
//last name
//favorite ice cream flavors
//Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores
//the favorite flavors.

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
	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for _, v := range p1.iceCream {
		fmt.Println(v)
	}
	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for _, v := range p2.iceCream {
		fmt.Println(v)
	}

}
