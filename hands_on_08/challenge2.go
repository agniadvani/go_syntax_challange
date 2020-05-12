//create a type person struct
//attach a method speak to type person using a pointer receiver
//*person
//create a type human interface
//to implicitly implement the interface, a human must have the speak method
//create func “saySomething”
//have it take in a human as a parameter
//have it call the speak method
//show the following in your code
//you CAN pass a value of type *person into saySomething
//you CANNOT pass a value of type person into saySomething
//here is a hint if you need some help

package main

import "fmt"

type person struct {
	name string
}
type human interface {
	speak()
}

func main() {
	p1 := person{
		"Micheal Scott",
	}
	saysomething(&p1)
	//saysomething(p1) does not run as the func speak recieves pointer to a person and not the type person.
}

func (p *person) speak() {
	fmt.Println(p.name)
}
func saysomething(h human) {
	h.speak()
}
