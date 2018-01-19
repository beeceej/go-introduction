# Example 3 | Interfaces!

* `interface {}` is the interface{} which all types belong to
* interfaces  allow you to switch out like functionality
* You DO NOT need to explicit define that a type implements an interface
  * All you need to do is implement the method signatures defined in the interface

```
type Adder interface {
	Add() interface{}
}
```

For a type to implement the above you must define a method with the signature

Types may `implement` many interfaces, and you don't need to define which it implements