package main

import (
	"fmt"
	"net/http"
	"os"
)

var PersonMap map[string]Person

func main() {

	PersonMap = make(map[string]Person)
	http.HandleFunc("/people/", PersonsRouter)
	http.HandleFunc("/people", PersonsRouter)
	http.HandleFunc("/", RootHandler)
	err := http.ListenAndServe("localhost:11111", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
