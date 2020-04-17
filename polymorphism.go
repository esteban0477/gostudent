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

type hotdog int

func main() {
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			"Miss",
			"Moneypenny",
		},
		ltk: true,
	}

	p1 := person{
		first: "Juan",
		last:  "Cano",
	}

	var y hotdog

	fmt.Printf("%T\n", sa1)
	sa1.speak()
	sa2.speak()
	fmt.Println()
	humanReceiver(sa1)
	humanReceiver(sa2)
	humanReceiver(p1)
	y.speak()
	humanReceiver(y)
}

func (p person) speak() {
	fmt.Println("Inside speak funtion from person struct")
}

func (s secretAgent) speak() {
	fmt.Println("Inside speak funtion from secretAgent struct")
}

func (hot hotdog) speak() {
	fmt.Println("I'm Calling my HotDog!")
}

func humanReceiver(h human) {
	switch h.(type) {
	case secretAgent:
		fmt.Println("I'm on human but my really type is secretAgent")
	case person:
		fmt.Println("I'm on human but my really type is person")
	case hotdog:
		fmt.Println("I'm on human but my really type is hotdog")
		fmt.Println(h.(hotdog))
	}
}
