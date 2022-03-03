package main

import "fmt"
import "encoding/json"
import "log"

type ContactInfo struct {
	Email   string `json:"email"`
	ZipCode int `json:"zipCode"`
}

type Person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	ContactInfo
}

func main() {
	jim := Person{
		FirstName: "Jim",
		LastName:  "Party",
		ContactInfo: ContactInfo{
			Email:   "jim@gmail.com",
			ZipCode: 94000,
		},
	}

	jim.updateName("jimmy")
	jim.print()
}

func (p *Person) updateName(newFirstName string) {
	(*p).FirstName = newFirstName
}

func (p *Person) print() {
	fmt.Printf("%%+v %+v \n", p)
	fmt.Printf("%%#v %#v \n", p)

	// Marshal
	pJSON, err := json.Marshal(p)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("Marshal %s:\n", string(pJSON))

	// MarshalIndent
	pJSON, err = json.MarshalIndent(p, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("MarshalIndent %s:\n", string(pJSON))
}
