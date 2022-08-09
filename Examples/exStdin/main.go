package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//Scanln

	var text string

	fmt.Print("Enter text: ")
	_, err := fmt.Scanln(&text)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(text)

	//Reader

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(text)

}
