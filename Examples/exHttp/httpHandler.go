package main

import (
	"fmt"
	"log"
	"net/http"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Main page ")
	w.Write([]byte("!!!"))

}

func pageHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Single page: ", r.URL.String())

}

func pagesHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Multiple pages: ", r.URL.String())

}

func main() {

	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/page", pageHandler)
	http.HandleFunc("/pages/", pagesHandler)
	log.Println("Starting server")
	http.ListenAndServe(":8080", nil)

}
