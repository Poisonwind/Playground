package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	log.Println("Server started")

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Nothing here")
		data, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(rw, "Something wrong", http.StatusBadRequest)
			return
		}

		log.Printf("Data: %s", data)
		fmt.Fprintf(rw, "Hello! You said: %s", data)
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello world")
	})
	http.HandleFunc("/bye", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Goodbye world")
	})

	http.ListenAndServe(":9090", nil)

}