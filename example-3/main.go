package main

import "fmt"


type Adder interface {
	Add() interface{}
}

type AddIntStruct struct {
	A int
	B int
}

type AddStringStruct struct {
	A string
	B string
}

func (a *AddIntStruct) Add() interface{} {
	return a.A + a.B
}

func (a *AddStringStruct) Add() interface{} {
	return a.A + a.B
}

func main() {

	adderIntOne()
	adderIntTwo()

	adderStringOne()
	adderStringTwo()

	adders := make([]Adder, 2, 2)

	adders[0] = &AddIntStruct{}
	adders[1] = &AddStringStruct{}
}

func adderIntOne() {
	// Create a new reference to the AddStruct and assign the pointer to adder
	adder := new(AddIntStruct)

	adder.A = 1
	adder.B = 2

	fmt.Println(adder.Add())
}

func adderIntTwo() {
	// Or create reference and assign values to the struct
	adder := &AddIntStruct{
		A: 1,
		B: 2,
	}

	fmt.Println(adder.Add())
}

func adderStringOne() {
	// Create a new reference to the AddStruct and assign the pointer to adder
	adder := new(AddStringStruct)

	adder.A = "Hello! "
	adder.B = "World! "

	fmt.Println(adder.Add())
}

func adderStringTwo() {
	// Or create reference and assign values to the struct
	adder := &AddStringStruct{
		A: "Hello! ",
		B: "World! ",
	}

	fmt.Println(adder.Add())
}
