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

	intAdder()
	stringAdder()

	adders := make([]Adder, 2, 2)

	adders[0] = &AddIntStruct{A: 1, B: 100}
	adders[1] = &AddStringStruct{A: "a", B: "b"}

	for _, v := range adders {
		fmt.Println(v.Add())
	}
}

func intAdder() {
	// Create a new reference to the AddStruct and assign the pointer to adder
	adder := new(AddIntStruct)

	adder.A = 1
	adder.B = 2

	fmt.Println(adder.Add())
}

func stringAdder() {
	// Create a new reference to the AddStruct and assign the pointer to adder
	adder := new(AddStringStruct)

	adder.A = "Hello! "
	adder.B = "World! "

	fmt.Println(adder.Add())
}
