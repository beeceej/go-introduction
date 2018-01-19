package coolpackage



// unexported :( so we can't access from outside of this package
type unCoolStruct struct {
}

// NewUncoolStruct , Unless we have an eported function which returns the struct
func NewUncoolStruct() unCoolStruct {
	return unCoolStruct{}
}
