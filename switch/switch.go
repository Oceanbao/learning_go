package main

import (
	"fmt"
)

func main() {
	switch ch := "b"; ch {
	case "a":
		fmt.Println("a")
	case "b", "c":
		fmt.Println("b or c")
	default:
		fmt.Println("No matching char")
	}

	age := 45
	switch {
	case age < 18:
		fmt.Println("Kid")
	case age >= 18 && age < 40:
		fmt.Println("Young")
	default:
		fmt.Println("Old")
	}

	i := 45
	switch {
	case i < 10:
		fmt.Println("i < 10")
		fallthrough
	case i < 50:
		fmt.Println("i < 50")
		fallthrough
	case i < 100:
		fmt.Println("i < 100")
	}

	// type switch
	printType("test_string")
}

func printType(t interface{}) {
	switch v := t.(type) {
	case string:
		fmt.Println("Type: string")
	case int:
		fmt.Println("Type: int")
	default:
		fmt.Println("Unknown Type %T", v)
	}
}
