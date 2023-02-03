package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

const (
	apiURL = "34.77.36.161"
)

func main() {
	var wg sync.WaitGroup
	resultChan := make(chan string, 1)

	fmt.Println("Let's go API")
	for port := 3000; port < 4001; port++ {
		wg.Add(1)
		go pingUrl(port, resultChan, &wg)
	}
	wg.Wait()
	close(resultChan)

	fmt.Println(<-resultChan)
}
func pingUrl(port int, resultChan chan<- string, wg *sync.WaitGroup) {
	url := "http://" + apiURL + ":" + strconv.Itoa(port)
	resp, err := http.Get(url)
	if err != nil {
		// fmt.Printf("Error GETting %d: %v\n", port, err)
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
	resultChan <- secretKey
	wg.Done()
}
