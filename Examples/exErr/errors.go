package main

import (
	"errors"
	"fmt"
)

var ErrSomethingBad = errors.New("something bad happen")

func PositiveNum(num int) (bool, error) {

	if num < 0 {
		return false, fmt.Errorf("no no no %w", ErrSomethingBad)
	}

	return true, nil
	
}

func main() {

	_, err := PositiveNum(-2)
	fmt.Println(err)

}