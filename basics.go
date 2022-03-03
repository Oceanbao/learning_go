package main

import (
	"fmt"
)


// Struct
type person struct {
	name string
	age int
}

func newPerson(name string, age int) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

// func impl on struct type *person (restricted?)
func (pp *person) changeName(newName string) {
	// (*pp).name = newName
	pp.name = newName
}

func printPointer(ptr *string) {
	fmt.Printf("&ptr: %p\n", &ptr)
}

func main() {
	// Slice
	zero2nine := []int{}
	for i := 0; i < 10; i++ {
		zero2nine = append(zero2nine, i)
	}

	fmt.Println("Full slice: ", zero2nine)

	for i, v := range zero2nine {
		if v%2 == 0 {
			fmt.Println(i, " is ", v, " -> even")
		} else {
			fmt.Println(i, " is ", v, " -> odd")
		}
	}

	// Map
	names := map[string]string{}
	names["Ocean"] = "male"

	fmt.Println(names)

	newNames := make(map[string]int, 5)
	newNames["Ocean"] = 33
	delete(newNames, "Ocean")

	fmt.Println(newNames)

	// Struct - see above CANNOT declare func inside main??
	s := person{name: "Bao", age: 33}
	sp := &s // ??
	sp.age = 99
	fmt.Printf("s: %v\n", s)
	fmt.Printf("s.name: %v\n", s.name)
	// fmt.Printf("(*s).name: %v\n", (*s).name) ERR: invalid redirect
	fmt.Printf("sp: %p\n", sp)
	fmt.Printf("sp.name: %v\n", sp.name)
	fmt.Printf("(*sp).name: %v\n", (*sp).name)
	fmt.Printf("&s: %p\n", &s)
	fmt.Printf("&sp: %p\n", &sp)
	fmt.Printf("(*(&sp)): %p\n", (*(&sp)))

	// seems even though func signature defined arg as
	// *person (pointer of person); it auto-cast the
	// person arg as *person!
	s.changeName("new name via person s")
	fmt.Println("s.changeNew() -> ", s)
	sp.changeName("new name via *person sp")
	fmt.Println("sp.changeNew() -> ", s)

	// NOTE: Go pass EVERYTHING by value!
	someString := "This is a string literal"
	ptrString := &someString

	fmt.Printf("ptrString: %p\n", ptrString)
	fmt.Printf("&ptrString: %p\n", &ptrString)
	// pass by value means func COPY ptrString's pointer
	printPointer(ptrString)


}


