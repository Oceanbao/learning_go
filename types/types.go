package main

import (
	"fmt"
	"math/bits"
	"reflect"
	"unsafe"
)

func main() {
	// const uintSize = 32 << (^unit(0) >> 32 & 1) // 32 or 64
	sizeOfIntInBits := bits.UintSize
	fmt.Printf("%d bits\n", sizeOfIntInBits)

	var a int
	fmt.Printf("%d bytes\n", unsafe.Sizeof(a))
	fmt.Printf("a type is %s\n", reflect.TypeOf(a))

	type sample struct {
		a int
		b string
	}

	s := &sample{a: 1, b: "test"}

	// Addr of field b in struct s
	p := unsafe.Pointer(uintptr(unsafe.Pointer(s)) + unsafe.Offsetof(s.b))

	// Typecasting it to string pointer and pprint value
	fmt.Println(*p)
	fmt.Println(*(*string)(p))

	var r byte = 'a'

	fmt.Printf("Size: %d\n", unsafe.Sizeof(r))
	fmt.Printf("Type: %s\n", reflect.TypeOf(r))
	fmt.Printf("Char: %c\n", r)
	fmt.Println([]byte("abc")) // decimal value of byte

	fmt.Printf("%U\n", []rune("0b$")) // unicode CodePoint
	fmt.Println([]rune("0b$")) // decoimal value of Unicode CodePoint

	fmt.Println([]byte("ab$")) // convert string to byte array (decimal)
	fmt.Println("len(0b$): ", len("0b$")) // 4, for 4 bytes

	fmt.Println("this\nthat") // double quote ensures escaping
	fmt.Println(`this\nthat`) // backtick string as is

}
