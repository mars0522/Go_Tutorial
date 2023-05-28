package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	myUrl := "http://localhost:2000/post"

	// JSON payload
	requestbody := strings.NewReader(`{"Name": "Varun", "Age": "25"}`)

	resp, err := http.Post(myUrl, "application/json", requestbody)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
