package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)


func main() {


	client := &http.Client{}
	req, err := http.NewRequest("GET","https://fxhapi.feixiaohao.com/public/v1/ticker", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := url.Values{}
	q.Add("start", "5")
	q.Add("limit", "5")
	q.Add("convert", "USD")

	req.Header.Set("Accepts", "application/json")
	req.URL.RawQuery = q.Encode()


	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request to server")
		os.Exit(1)
	}
	fmt.Println(resp.Status)
	respBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(respBody))
}
