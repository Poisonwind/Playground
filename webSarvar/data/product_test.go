package data

import "testing"

func TestCheckValidationWorks(t *testing.T) {
	prod := &Product{
		Name: "tea",
		Price: 1.50,
		SKU: "abs-ddtf-askju",
	}
	err := prod.Validate()

	if err != nil {
		t.Error(err)
	}
}