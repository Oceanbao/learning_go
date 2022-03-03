package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

	// -----------
	var h human // explicit declaring human interface type

	// this error as declaring value receiver
	// h = employee{name: "John"} // assign employee struct as human type
	h = &employee{name: "John"} // assign employee struct as human type
	h.breathe()
	h.walk()
	h.speak()

	d := dog{age: 2}
	p1 := pet1{name: "Pet1", animal: d}
	p1.breathe()
	p2 := pet2{name: "Pet2", a: d}
	p2.a.breathe()

	printType := func (a animal) {
		l := a.(lion) // assert a to be lion type
		fmt.Printf("Age: %d\n", l.age)
		// further assert a.(dog) panics

		d, ok := a.(dog)
		if ok {
			fmt.Println(d)
		} else {
			fmt.Println("a is not of type dog")
		}
	}

	// printType(h) panic, h is *main.employee not main.lion
	// printType(d) panic, d is main.dog not main.lion
	// printType(p1) panic, p1 is main.pet1 not main.lion
	l := lion{age: 3}
	printType(l)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// VERY custom logic for generating an english greeting
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

// embed interface in inerface
type animal interface {
	breathe()
	walk()
}

type lion struct {
	age int
}

func (d lion) breathe() {
	fmt.Println("Lion breathes")
}

func (d lion) walk() {
	fmt.Println("Lion walks")
}

type dog struct {
	age int
}

func (d dog) breathe() {
	fmt.Println("Dog breathes")
}

func (d dog) walk() {
	fmt.Println("Dog walks")
}

// interface type human is valid iif impl all method in human recursively
type human interface {
	animal
	speak()
}

type employee struct {
	name string
}

func (e *employee) breathe() {
	fmt.Println("Employee breathes")
}

func (e *employee) walk() {
	fmt.Println("Employee walks")
}

func (e *employee) speak() {
	fmt.Println("Employee speaks")
}

// named vs unnamed interface
type pet1 struct {
	animal // unnamed, thus its method accessible directly
	name string
}
type pet2 struct {
	a animal // named, thus its method accessible via a only
	name string
}
