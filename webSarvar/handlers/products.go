package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"websarvar/data"
)

type Product struct {
	log *log.Logger
}

func NewProduct (l *log.Logger) *Product {
	return &Product{l}
}

func (p *Product) ServeHTTP(rw http.ResponseWriter, r *http.Request){

	productList := data.GetProducts()
	data, err := json.Marshal(productList)
	if err != nil {
		http.Error(rw, "unable to marshal json", http.StatusInternalServerError)
	}

	rw.Write(data)

}