package main

import (
	"encoding/json"
	"fmt"
)

// Product isStruct for serialization, note the json tags, which influence the json produced
type Product struct {
	ProductID         int      `json:"productId"`
	OrganizationGroup string   `json:"organizationGroup"`
	Ingredients       []string `json:"ingredients"`
}

func main() {
	p := &Product{
		ProductID:         1,
		OrganizationGroup: "letsLearnGo!",
		Ingredients:       []string{"salt", "water", "glucose"},
	}

	b, err := marshal(p)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(b))

	newProduct := unMarshal(b)

	fmt.Println(newProduct)
}

func marshal(p *Product) ([]byte, error) {
	return json.MarshalIndent(p, "", "  ")
}

func unMarshal(b []byte) *Product {
	p := &Product{}
	_ = json.Unmarshal(b, p)
	return p
}
