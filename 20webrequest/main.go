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

	bodyinbytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	content := string(bodyinbytes)
	fmt.Println(content)
}
