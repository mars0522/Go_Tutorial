package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, err := http.Get("https://github.com/")

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	// Reading Body from response which contains string  content in Bytes format
	bodyinbytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	// Converting the bytes into a string
	content := string(bodyinbytes)
	fmt.Println(content)
}
