package main

import (
	"log"
	"net/http"
	"os"
	"websarvar/handlers"
)

func main() {	
	myLog := log.New(os.Stdout, "webSarvar ", log.LstdFlags)
	
	myServeMux := http.NewServeMux()
	helloHandler := handlers.NewHello(myLog)
	goodByeHandler := handlers.NewGoodbye(myLog)

	myServeMux.Handle("/", helloHandler)
	myServeMux.Handle("/gb/", goodByeHandler)

	myLog.Println("Server started")
	http.ListenAndServe(":9090", myServeMux)	

}