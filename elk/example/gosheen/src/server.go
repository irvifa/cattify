package main

import (
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	port := os.Getenv("PORT")

	if port == "" {
		port = "8082"
	}

	http.HandleFunc("/", Check)
	Info.Println("listening at " + port + "... ")
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		Error.Println(err)
	}
}
