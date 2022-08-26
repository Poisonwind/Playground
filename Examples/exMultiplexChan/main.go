package main

import (
	"fmt"
	"sync"
)

func carsNames() <-chan string {

	c := make(chan string)
	wg := sync.WaitGroup{}
	items := []string{"mazda", "honda", "suzuki"}

	wg.Add(1)
	go func(){
		defer wg.Done()
		for _, val := range items {
			c <- val
		}
	}()

	go func(){
		wg.Wait()
		close(c)
	}()	

	return c
}

func numsNames() <-chan string {

	c := make(chan string)
	wg := sync.WaitGroup{}
	items := []string{"one", "two", "three"}

	wg.Add(1)
	go func(){
		defer wg.Done()
		for _, val := range items {
			c <- val
		}
	}()

	go func(){
		wg.Wait()
		close(c)
	}()	

	return c
}

func formsNames() <-chan string {

	c := make(chan string)
	wg := sync.WaitGroup{}
	items := []string{"circle", "square", "triangle"}

	wg.Add(1)
	go func(){
		defer wg.Done()
		for _, val := range items {
			c <- val
		}
	}()

	go func(){
		wg.Wait()
		close(c)
	}()	

	return c
}

func multiplexingFunc(channels ...<-chan string) <-chan string {

	multiplexedChan := make(chan string)
	wg := sync.WaitGroup{}

	wg.Add(len(channels))
	for _, c := range channels {
		go func(c <-chan string) {
			defer wg.Done()
			for i := range c {
				multiplexedChan <- i 
			}
		}(c)
	}

	go func(){
		wg.Wait()
		close(multiplexedChan)
	}()


	return multiplexedChan
}





func main() {

	channels := []<-chan string{}
	channels = append(channels, carsNames(), formsNames(), numsNames())

	multiChan := multiplexingFunc(channels...)

	for data := range multiChan {
		fmt.Println(data)
	}

}
