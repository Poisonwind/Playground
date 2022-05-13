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

func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

func (p *Product) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProducts(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		p.log.Println("Use PUT Method", r.URL.Path)

		exp := regexp.MustCompile(`/(\d+)`)
		subStrings := exp.FindAllStringSubmatch(r.URL.Path, -1)

		if len(subStrings) != 1 {
			p.log.Println("Ivalid URL more then one id")
			http.Error(rw, "Invalid URL", http.StatusBadRequest)
		}

		if len(subStrings[0]) != 2 {
			p.log.Println("Ivalid URL more then one regexp group")
			http.Error(rw, "Invalid URL", http.StatusBadRequest)
		}

		idString := subStrings[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			p.log.Println("Ivalid URL invalid to convert to number", idString)
			http.Error(rw, "Invalid URL", http.StatusBadRequest)
		}
		p.log.Println("got id", id)
		p.updateProducts(id, rw, r)
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
	data.AddProduct(prod)
}

func (p *Product) updateProducts(id int, rw http.ResponseWriter, r *http.Request) {

	p.log.Println("Handle PUT Products")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "unable to unmarshal json", http.StatusBadRequest)
	}

	p.log.Printf("new prod: %#v", prod)
	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "product not found", http.StatusInternalServerError)
	}
}
