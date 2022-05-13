package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"
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

	if r.Method == http.MethodPut {
		exp := regexp.MustCompile(`/(\d+)`)
		subStrings := exp.FindAllStringSubmatch(r.URL.Path, -1)

		if len(subStrings) != 1 {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
		}

		if len(subStrings[0]) != 1 {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
		}

		idString := subStrings[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			p.log.Println(err)
		}
		p.log.Println("got id", id)
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
	data.AddProduct(prod)
}