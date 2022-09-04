package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "Coffee",
		Price: 1.0,
		SKU:   "a-a-a",
	}

	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
