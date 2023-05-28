package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {

	myurl := "http://localhost:2000/postform"
	data := url.Values{}
	data.Add("name", "Varun")
	data.Add("email", "anpch@example.com")
	data.Add("password", "123456")

	res, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("OK", string(body))
}
