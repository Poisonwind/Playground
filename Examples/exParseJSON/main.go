package main

import (
	"encoding/json"
	"fmt"
)

type jsonHere struct {
	Num    int    `json:"num"` //12324
	Number int    `json:"number"`
	Text   string `json:"text"`
	Txt    string `json:"txt"` //ывавы
}

func main() {

	j := jsonHere{}

	jsonExample := `
	{
		"num":15,
		"number": 223,
		"text":"hello",
		"txt":"woot"
	}`

	fmt.Println(jsonExample)

	err := json.Unmarshal([]byte(jsonExample), &j)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v", j)

}
