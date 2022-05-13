package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product

func (p *Products) ToJSON(writer io.Writer) error {

	encoder := json.NewEncoder(writer)
	return encoder.Encode(p)

}

func (p *Product) FromJSON(reader io.Reader) error {

	decoder := json.NewDecoder(reader)
	return decoder.Decode(p)

}

func GetProducts() Products {
	return ProductList
}

//get id for new product
func getNextID() (newId int) {
	lastProduct := ProductList[len(ProductList)-1]
	newId = lastProduct.ID + 1
	return
}

//add new product to product list
func AddProduct(p *Product) {
	p.ID = getNextID()
	ProductList = append(ProductList, p)
}

func UpdateProduct(id int, p *Product) error {

	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}

	p.ID = id
	ProductList[pos] = p
	return nil
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error) {
	for pos, prod := range ProductList {
		if prod.ID == id {
			return prod, pos, nil
		}
	}

	return nil, -1, ErrProductNotFound
}

var ProductList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky cofee",
		Price:       2.45,
		SKU:         "abc123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffe without milk",
		Price:       1.99,
		SKU:         "asd321",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
