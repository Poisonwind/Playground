package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"websarvar/data"

	"github.com/gorilla/mux"
)

type Product struct {
	log *log.Logger
}

func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}



func (p *Product) GetProducts(rw http.ResponseWriter, r *http.Request) {

	p.log.Println("Handle GET Products")

	productList := data.GetProducts()
	err := productList.ToJSON(rw)

	if err != nil {
		http.Error(rw, "unable to marshal json", http.StatusInternalServerError)
	}

}

func (p *Product) AddProducts(rw http.ResponseWriter, r *http.Request) {

	p.log.Println("Handle POST Products")

	prod := r.Context().Value(ProductKey{}).(*data.Product)

	p.log.Printf("new prod: %#v", prod)
	data.AddProduct(prod)
}

func (p *Product) UpdateProducts(rw http.ResponseWriter, r *http.Request) {

	vals := mux.Vars(r)
	id, err := strconv.Atoi(vals["id"])
	if err != nil {
		http.Error(rw, "unable to convert id", http.StatusBadRequest)
		return
	}


	p.log.Println("Handle PUT Products", id)

	prod := r.Context().Value(ProductKey{}).(*data.Product) // .(*data.Product) means type of prod for next functions

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

type ProductKey struct{}

func (p *Product) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request){
		//parsing json
		prod := &data.Product{}
		err := prod.FromJSON(r.Body)
	
		if err != nil {
			p.log.Println("[ERROR] deserializing product", err)
			http.Error(rw, "error reading product", http.StatusBadRequest)
			return
		}

		//validate &Product
		err = prod.Validate()
		if err != nil {
			p.log.Println("[ERROR] validating product", err)
			http.Error(
				rw,
				fmt.Sprintf("error validating product: %s", err),
				http.StatusBadRequest)
			return			
		}

		//add product to context for passing through
		ctx := context.WithValue(r.Context(), ProductKey{}, prod)
		req := r.WithContext(ctx)
		
		//call next handler
		next.ServeHTTP(rw, req)
	})

	
}
