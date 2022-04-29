package main

import (
	"log"
	"net/http"
	"os"
	"websarvar/handlers"
)

func main() {	
	myLog := log.New(os.Stdout, "webSarvar", log.LstdFlags)
	
	myServeMux := http.NewServeMux()
	helloHandler := handlers.NewGreet(myLog)

	myServeMux.Handle("/", helloHandler)

	log.Println("Server started")
	http.ListenAndServe(":9090", myServeMux)
	
	

}