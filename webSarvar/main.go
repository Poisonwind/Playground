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

	"github.com/gorilla/mux"
)

func main() {	

	//INIT//

	myLog := log.New(os.Stdout, "webSarvar ", log.LstdFlags)

	myRouter := mux.NewRouter()

	getRouter := myRouter.Methods(http.MethodGet).Subrouter()
	putRouter := myRouter.Methods(http.MethodPut).Subrouter()
	postRouter := myRouter.Methods(http.MethodPost).Subrouter()

	myServer := &http.Server{
		Addr: ":9090",
		Handler: myRouter,
		IdleTimeout: 120*time.Second,
		ReadTimeout: 1*time.Second,
		WriteTimeout: 1*time.Second,
	} 

	//HANDLERS//

	helloHandler := handlers.NewHello(myLog)
	productHandler := handlers.NewProduct(myLog)
	goodByeHandler := handlers.NewGoodbye(myLog)

	//ROUTERS//

	//GET//
	getRouter.HandleFunc("/", helloHandler.GetJustHello)
	getRouter.HandleFunc("/gb/", goodByeHandler.GetJustGoodbye)
	getRouter.HandleFunc("/products/", productHandler.GetProducts)

	//POST//
	postRouter.Use(productHandler.MiddlewareProductValidation)
	postRouter.HandleFunc("/products/", productHandler.AddProducts)
	

	//PUT//
	putRouter.Use(productHandler.MiddlewareProductValidation)
	putRouter.HandleFunc("/products/{id:[0-9]+}", productHandler.UpdateProducts)
	

	// START //

	go func() {
		err := myServer.ListenAndServe()
		if err != nil {
			myLog.Fatal(err)
		}	
	}()	

	myLog.Println("Server started")
	
	// SHUTDOWN //

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	sig := <-sigChan
	myLog.Println("Graceful shutdown", sig)

	myContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	myServer.Shutdown(myContext)
}