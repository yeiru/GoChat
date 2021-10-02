package main

import (
	"fmt"
	"net/http"

	"github.com/yeiru/GoChat/helpers"
)

func main() {
	// fmt.Println("Hello World")
	var myVariable helpers.MyStruct
	myVariable.MyString = "Testing"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello world "+myVariable.MyString)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("number of bytes written ", n)
	})

	http.ListenAndServe(":8080", nil)
}
