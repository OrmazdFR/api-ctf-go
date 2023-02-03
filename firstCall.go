package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

func getFirstSecret(port int, resultChan chan<- string, wg *sync.WaitGroup) {
	url := "http://" + apiURL + ":" + strconv.Itoa(port)
	resp, err := http.Get(url)
	if err != nil {
		wg.Done()
		return
	}
	if resp.Body == nil {
		wg.Done()
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var secretKey = strings.Split(string(body), "The secret key is: ")[1]
	resultChan <- strings.Trim(secretKey, " ")
	wg.Done()
}
