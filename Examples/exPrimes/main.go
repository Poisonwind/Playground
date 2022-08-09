package main

import "fmt"

func primeNumber(n int) {

	var (
		counter int
		prime bool 
	)

	for i := 2; counter < n; i++ {
		prime = true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				prime = false
			}
		}
		if prime {
			fmt.Println(i)
			counter++
		}
		
	}

}

func main() {

	primeNumber(7)

}