package main

import "fmt"

const a = "THIS IS A CONSTANT"

// an array with size 1
var anArray = [1]string{"1"}

//make a slice with capacity and size of 1
var aSlice = make([]string, 1)

//make a slice with capacity 1 and size 10
var aSlice2 = make([]string, 1, 10)

// structs are kind of `like classes in Java`
type myStruct struct {
	a int
	b string
	c float64
	d []byte
}

// the variable "aStruct" is of type `myStruct`
var aStruct myStruct

// functions are first class
var aFunction = func() string { return "I'M A FUNCTION" }

// See?!
func testMyFunction(f mySuperSpecialFunction) int {
	return f("")
}

// we can have function types as well
type mySuperSpecialFunction func(a string) int

func main() {
	fmt.Println(`
		const a = "THIS IS A CONSTANT"

		// an array with size 1
		var anArray = [1]string{"1"}
		
		//make a slice with capacity and size of 1
		var aSlice = make([]string, 1)
		
		//make a slice with capacity 1 and size 10
		var aSlice2 = make([]string, 1, 10)
		
		// structs are kind of like classes in Java
		type myStruct struct {
			a int
			b string
			c float64
			d []byte
		}
		
		// the variable "aStruct" is of type myStruct
		var aStruct myStruct
		
		// functions are first class
		var aFunction = func() string { return "I'M A FUNCTION" }
		
		// See?!
		func testMyFunction(f mySuperSpecialFunction) int {
			return f("")
		}
		
		// we can have function types as well
		type mySuperSpecialFunction func(a string) int)`)
}
