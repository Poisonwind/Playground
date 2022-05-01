package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	log *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (hello *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(rw, "Something wrong", http.StatusBadRequest)
		return
	}

	if string(data)== "" {
		fmt.Fprintln(rw, "Hello! You said nothing!")
		hello.log.Println("There is no data")
	} else {
		fmt.Fprintf(rw, "Hello! You said: %s", data)
		hello.log.Printf("Data: %s", data)
	}
	

}