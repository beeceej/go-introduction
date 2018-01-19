package main

import (
	"fmt"

	"github.com/beeceej/go-introduction/example-2/coolpackage"
)

func main() {

	// We can access the cool values but not the uncool values
	cool := coolpackage.CoolStruct{CoolValue: "A Cool Value"}
	fmt.Println(cool.CoolValue)

	// We can call methods on a struct like
	fmt.Println(cool.DoSomethingCool())

	// we cant access uncoolStruct directly, but we can call the function which returns it
	uncool := coolpackage.NewUncoolStruct()
	fmt.Println(uncool)

	// In general the casing of the first letter determines package visibility
}
