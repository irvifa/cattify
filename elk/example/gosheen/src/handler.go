package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

// curl -X POST -H "Content-Type: application/json" -d '{"asked":"1"}' http://127.0.0.1:8082

func Check(res http.ResponseWriter, req *http.Request) {
	Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	Info.Println("inside the handler..")
	var msg Asked
	test := json.NewDecoder(req.Body)
	err := test.Decode(&msg)

	if err != nil {
		Error.Println(err)
	}

	num := msg.Num
	is_num := is_number(num)

	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(res).Encode(is_num); err != nil {
		Error.Println(err)
	}
}
