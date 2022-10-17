package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Ping(url string) string {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
	}

	client := &http.Client{Timeout: time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}

	answer, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	return string(answer)

}

func main() {

}
