//Creating interfaces.
package main

import "fmt"

type person struct {
	first string
	last  string
}
type secretAgent struct {
	person
	ltk bool
}
type human interface {
	speak()
}

func main() {
	p1 := person{
		first: "Pam",
		last:  "Beasely",
	}
	sa1 := secretAgent{
		person: person{
			first: "Micheal",
			last:  "Scott",
		},
		ltk: false,
	}
	sa2 := secretAgent{
		person: person{
			first: "dwight",
			last:  "shrute",
		},
		ltk: true,
	}
	sa1.speak()
	sa2.speak()
	p1.speak()
	fmt.Println("-------------")
	humanCheck(sa1)
	humanCheck(sa2)
	humanCheck(p1)
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, "- a secret agent")
}
func (p person) speak() {
	fmt.Println("I am", p.first, p.last, "- a person")
}
func humanCheck(h human) {
	switch h.(type) {
	case person:
		fmt.Println("This is a proof that", h.(person).first, "is a human.")
	case secretAgent:
		fmt.Println("This is a proof that", h.(secretAgent).first, "is a human.")
	}
}
