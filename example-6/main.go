package main

import (
	"encoding/json"
	"net/http"
)

type Product struct {
	ProductID         int      `json:"productId"`
	OrganizationGroup string   `json:"organizationGroup"`
	Ingredients       []string `json:"ingredients"`
}

func main() {
	http.HandleFunc("/products", func(rw http.ResponseWriter, r *http.Request) {
		p := &Product{
			ProductID:         1,
			OrganizationGroup: "letsLearnGo!",
			Ingredients:       []string{"salt", "water", "glucose"},
		}
		b, err := json.MarshalIndent(p, "", "  ")
		if err != nil {
			rw.Write([]byte("Houston we got a problem: " + err.Error()))
		}
		rw.Write(b)
	})
	http.ListenAndServe(":8080", nil)
}
