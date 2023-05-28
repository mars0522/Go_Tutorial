package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"couse_name"`
	Price    int      `json:"price"`
	Platfrom string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {

	offeredCourses := []course{
		{"Ful Stack", 299, "www.fulstack.com", "123@fdf", []string{"WebDevelopment"}},
		{"React", 199, "www.fulstack.com", "123@fdf", []string{"React"}},
		{"GoLang", 599, "www.fulstack.com", "123@fdf", []string{"Golang"}},
		{"Python", 299, "www.fulstack.com", "123@fdf", nil},
	}

	// Package this data as JSON data
	jsonData, err := json.MarshalIndent(offeredCourses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonData))

}
