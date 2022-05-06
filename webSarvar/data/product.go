package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID          int `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       float32 `json:"price"`
	SKU         string `json:"sku"`
	CreatedOn   string `json:"-"`
	UpdatedOn   string `json:"-"`
	DeletedOn   string `json:"-"`
}

type Products []*Product

func (p *Products) ToJSON (writer io.Writer) error {

	encoder := json.NewEncoder(writer)
	return encoder.Encode(p)

}

func (p *Product) FromJSON (reader io.Reader) error {

	decoder := json.NewDecoder(reader)
	return decoder.Decode(p)

}

func GetProducts() Products {
	return ProductList
}

var ProductList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky cofee",
		Price:       2.45,
		SKU:         "abc123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
	&Product{
		ID: 2,
		Name: "Espresso",
		Description: "Short and strong coffe without milk",
		Price: 1.99,
		SKU: "asd321",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
}