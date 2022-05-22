package main

import (
	"context"
	"math/rand"
	"sync"
	"time"
)

func waitTaxi(ctx context.Context, taxiName string, out chan<- string, wg *sync.WaitGroup) {

	defer wg.Done()

	waitTime := time.Duration(rand.Intn(100)+100) * time.Millisecond	
	println("taxi", taxiName, "route time", waitTime)

	select {
	case <- ctx.Done():
		return
	case <- time.After(waitTime):
		println("taxi", taxiName, "is ready")
		out <- taxiName
	}


}

func main() {

	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())

	result := make(chan string, 1)
	taxiServices := []string{"YandexTaxi", "GetTaxi", "GoTaxi", "Uber"}

	for _, val := range(taxiServices){
		wg.Add(1)
		go waitTaxi(ctx, val, result, wg)
	}

	firstTaxi := <- result
	cancel()

	println("the first taxi is", firstTaxi)

}