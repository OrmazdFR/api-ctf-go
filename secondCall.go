package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func postSecondSecretAndGetLink(secretStr string) {
	myUrl := "http://" + apiURL + ":3941"
	data := url.Values{
		"secretKey": {secretStr},
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
	return
}
