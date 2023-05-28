package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	url := "http://localhost:2000/get"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Println("Status code:", resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println("OK", string(body))

	var resString strings.Builder

	bytecount, err := resString.Write(body)

	if err != nil {
		panic(err)
	}
	fmt.Println(bytecount)

	fmt.Println(resString.String())

}
