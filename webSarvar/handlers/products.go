package handlers

import (
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

	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProducts(rw, r)
		return
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Product) getProducts(rw http.ResponseWriter, r *http.Request) {

	p.log.Println("Handle GET Products")

	productList := data.GetProducts()
	err := productList.ToJSON(rw)

	if err != nil {
		http.Error(rw, "unable to marshal json", http.StatusInternalServerError)
	}
	
}

func (p *Product) addProducts(rw http.ResponseWriter, r *http.Request) {

	p.log.Println("Handle POST Products")
	
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "unable to unmarshal json", http.StatusBadRequest)
	}

	p.log.Printf("new prod: %#v", prod)
}