package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"websarvar/handlers"
)

func main() {	
	myLog := log.New(os.Stdout, "webSarvar ", log.LstdFlags)
	
	myServeMux := http.NewServeMux()
	helloHandler := handlers.NewHello(myLog)
	goodByeHandler := handlers.NewGoodbye(myLog)
	productHandler := handlers.NewProduct(myLog)

	myServeMux.Handle("/", helloHandler)
	myServeMux.Handle("/gb/", goodByeHandler)
	myServeMux.Handle("/product/", productHandler)

	myServer := &http.Server{
		Addr: ":9090",
		Handler: myServeMux,
		IdleTimeout: 120*time.Second,
		ReadTimeout: 1*time.Second,
		WriteTimeout: 1*time.Second,
	}

	go func() {
		err := myServer.ListenAndServe()
		if err != nil {
			myLog.Fatal(err)
		}	
	}()	

	myLog.Println("Server started")
	
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	sig := <-sigChan
	myLog.Println("Graceful shutdown", sig)

	myContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	myServer.Shutdown(myContext)

	defer cancel()

}