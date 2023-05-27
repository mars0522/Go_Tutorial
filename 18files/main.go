package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	content := "Hi there, this is the text that needs to be added to the file"
	file, err := CreateFile("test.txt")
	if err != nil {
		panic(err)
	}

	length, err := file.WriteString(content)
	if err != nil {
		panic(err)
	}

	fmt.Printf("File length: %d\n", length)
	defer file.Close()
	FileRead("./test.txt")

}

func CreateFile(filename string) (*os.File, error) {

	file, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func FileRead(filename string) {
	databytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	fmt.Println("Content of the file is given below \n", string(databytes))
}
