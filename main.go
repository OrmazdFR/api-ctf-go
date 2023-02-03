package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sync"
)

const (
	apiURL    = "34.77.36.161"
	finalPort = "3610"
	finalKey  = "8116fdd3f12b6d7c4b136cbdaa3360a57eb4eb676ae63294450ee1f4f34b36f3"
)

func main() {
	var wg sync.WaitGroup
	resultChan := make(chan string, 1)

	fmt.Println("Let's go API")
	for port := 3000; port < 4001; port++ {
		wg.Add(1)
		go getFirstSecret(port, resultChan, &wg)
	}

	wg.Wait()
	close(resultChan)

	var firstSecret = <-resultChan
	postSecondSecretAndGetLink(firstSecret)

	lastApi()
}

func lastApi() {
	myUrl := "http://" + apiURL + ":" + finalPort
	data := url.Values{
		"finalKey": {finalKey},
	}

	resp, err := http.PostForm(myUrl, data)
	if err != nil {
		fmt.Printf("Error POSTing : %v\n", err)
		return
	}
	if resp.Body == nil {
		fmt.Println("resp body nil")
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}
