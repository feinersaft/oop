package main

import "fmt"

// The basic structure of an object in Golang.
type Object struct {
	name string
}

// The struct Object has the method sagHallo(), which returns its name property.
func (o Object) sagHallo() {
	fmt.Println("Hallo, ich bin", o.name)
}

func main()  {
	// Create an instance of the object with its name property by using the shorthand variable notation.
	objOne := Object{"Mrs. One"}

	/* 
		Create an instance of the object with its name property by using explicit notation.
		It initializies the object with the default "zero value", when no input values are provided:
		false 	for booleans
		0		for numeric types
		""		for strings
		nil 	for pointers, functions, interfaces, slices, channels, and maps
	*/
	var objTwo Object // objTwo.name == ""
	objThree := new(Object) // does the same as above

	// Call the only method in this struct.
	objOne.sagHallo()
	objTwo.sagHallo()
	objThree.sagHallo()

	// We can assign values anytime later:
	objTwo.name = "Mr. Two"
	objThree.name = "Senor Three"

	objTwo.sagHallo()
	objThree.sagHallo()
}