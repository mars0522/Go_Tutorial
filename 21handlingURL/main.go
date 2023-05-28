package main

import (
	"fmt"
	"net/url"
)

const ur string = "http://lco.dev:3000/learn?coursename=reactjs&paymentid=2121h3232ew"

func main() {

	// Parsing the URL string
	res, err := url.Parse(ur)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Host)
	fmt.Println(res.Scheme)
	fmt.Println(res.Path)
	fmt.Println(res.Query())
	fmt.Println(res.Port())
	fmt.Println(res.RawQuery)

	queryParams := res.Query()

	for key, val := range queryParams {
		fmt.Println(key, val)
	}

	// Making the URL from the URL components

	u := &url.URL{
		Scheme:   "http",
		Host:     "lco.dev",
		Path:     "/learn",
		RawQuery: "coursename=reactjs&paymentid=2121h3232ew",
	}

	anotherURL := u.String()
	fmt.Println("Another URL is: ", anotherURL)
}
