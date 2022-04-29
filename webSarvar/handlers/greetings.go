package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Greet struct {
	log *log.Logger
}

func NewGreet(l *log.Logger) *Greet {
	return &Greet{l}
}

func (greet *Greet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	greet.log.Println("Nothing here")
	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(rw, "Something wrong", http.StatusBadRequest)
		return
	}

	log.Printf("Data: %s", data)
	
	if string(data)== "" {
		fmt.Fprintln(rw, "Hello! You said nothing!")
	} else {
		fmt.Fprintf(rw, "Hello! You said: %s", data)
	}
	

}