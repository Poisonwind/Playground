package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"` //sku - function for validation
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

//validation method for &Product
func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)

	return validate.Struct(p)
}

//validation func for SKU
func validateSKU(field validator.FieldLevel) bool {
	//format sku: abc-abcd-abcde
	re, err := regexp.Compile(`[a-z]+-[a-z]+-[a-z]+`)
	if err != nil {
		fmt.Errorf(err.Error())
	}

	matches := re.FindAllString(field.Field().String(), -1)

	return len(matches) == 1
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
