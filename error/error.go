package main

/*
* use type impl error interface - conventional way
* to repr error 
*
* use panic and recover
*/

/*
* type error interface {
	Error() string
	* }
*/

/*
* os.Open()
*
* func Open(name string) (*File, error)
* if error, -> *PathError impl error interface
*
* type PathError struct {
	Op string
	Path string
	Err error
* }
*
* pointer to PathError impl Error() hence impl error interface
*
* func (e *PathError) Error() string { return e.Op + " " + e.Path + ": " + e.Err.Error() }
*
* Hence error can be returned as second value from os.Open
*
* default 0 vaule of interface is nil and that error is interface
*
*/

import (
	"fmt"
	"os"
	"errors"
	"runtime/debug"
)

func main() {
	file, err := os.Open("non-existing.txt")
	if err != nil {
		// assert error type
		if e, ok := err.(*os.PathError); ok {
			fmt.Printf("Assert Error e is of type PathError. Path: %v\n", e.Path)
			fmt.Println("Error.message: ", err)
		}
	} else {
		fmt.Println(file.Name() + "opened ok")
	}

	var pathError *os.PathError
	if errors.As(err, &pathError) {
		fmt.Printf("Using errors.As to check type: %v\n", pathError.Path)
	}

	// manual error
	sampleErr1 := errors.New("error occured")
	fmt.Println(sampleErr1)

	sampleErr2 := fmt.Errorf("Err is: %s", "database conn issue")
	fmt.Println(sampleErr2)

	errInput := validate("", "")
	if errInput != nil {
		// type assertion
		if err, ok := errInput.(*inputError); ok {
			fmt.Println("ok: ", ok)
			fmt.Println(err) // fmt print will infer error type and print err.message
			// custom method of error type
			fmt.Printf("Missing Field is %s\n", err.getMissingField())
		}
		if _, ok := errInput.(errorOne); ok {
			fmt.Println("...")
		} else {
			fmt.Println("type asserting errInput.(errorOne): ok: ", ok)
		}
	}

	e1 := errorOne{}
	// %w in Errorf usd to wrap error
	e2 := fmt.Errorf("E2: %w", e1)
	e3 := fmt.Errorf("E3: %w", e2)
	fmt.Println(e2)
	fmt.Println(e3)

	num := 3
	err = checkPositiveAndEven(num)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Given num is positive and even")
	}

	// unwrap error
	fmt.Println(errors.Unwrap(e3))
	fmt.Println(errors.Unwrap(e2))

	// panic and recover
	var panicHandler = func() {
		if r := recover(); r != nil {
			fmt.Println("Recovering from panic:", r)
			// stack trace
			fmt.Println("Stack Trace:")
			debug.PrintStack()
		}
	}

	var checkAndGet = func(a []int, index int) (value int, err error) {
		// this is returned when panic-recover
		value = 10
		defer panicHandler()
		if index > (len(a) - 1) {
			panic("Out of bound access for slice")
		}
		// normal flow
		value = a[index]

		return value, nil
	}

	a := []int{5, 6}
	valLast, errLast := checkAndGet(a, 2)
	fmt.Printf("Val: %d\n", valLast)
	fmt.Println("Error: ", errLast)

}

type inputError struct {
	// message field is the interface field for error type
	message string // mandatory
	missingFiled string
}

func (i *inputError) Error() string {
	return i.message
}

func (i *inputError) getMissingField() string {
	return i.missingFiled
}

func validate(name, gender string) error {
	if name == "" {
		return &inputError{message: "Name is mandatory", missingFiled: "name"}
	}
	if gender == "" {
		return &inputError{message: "Gender is mandatory", missingFiled: "gender"}
	}
	return nil // important
}

type errorOne struct {}

func (e errorOne) Error() string {
	return "Error One occurred"
}

// wrapping error use case
type notPositive struct {
	num int
}

func (e notPositive) Error() string {
	return fmt.Sprintf("checkPositive: Given num %d not possitive.", e.num)
}

func checkPositive(num int) error {
	if num < 0 {
		return notPositive{num: num}
	}
	return nil
}

type notEven struct {
	num int
}

func (e notEven) Error() string {
	return fmt.Sprintf("checkEven: Given num %d is not even", e.num)
}

func checkEven(num int) error {
	if num % 2 == 1 {
		return notEven{num: num}
	}
	return nil
}

func checkPositiveAndEven(num int) error {
	if num > 100 {
		return fmt.Errorf("checkPositiveAndEven: Num is %d greater than 100", num)
	}

	err := checkPositive(num)
	if err != nil {
		return fmt.Errorf("checkPositiveAndEven: %w", err)
	}

	err = checkEven(num)
	if err != nil {
		return fmt.Errorf("checkPositiveAndEven: %w", err)
	}

	return nil
}
