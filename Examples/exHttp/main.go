package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func parseTags(url string)(err error){

	return
}

func sendRequest(url string)(err error){

	client := http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	bytesReaded, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bytesReaded))
	return nil
}

func main() {

	err := sendRequest("http://vpustotu.ru/moderation/")
	if err != nil {
		log.Fatal(err)
	}

}