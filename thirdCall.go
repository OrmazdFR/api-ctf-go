package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func thirdCall() {
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
