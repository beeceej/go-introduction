package coolpackage

// CoolStruct is Exported
type CoolStruct struct {
	CoolValue   string
	uncoolValue string
}

// DoSomethingCool is a method attached to the CoolStruct, and may only be called if we have an instance of CoolStruct
func (c *CoolStruct) DoSomethingCool() string {

	return c.CoolValue
}
