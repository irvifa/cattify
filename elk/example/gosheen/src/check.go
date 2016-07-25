package main

import (
	"io/ioutil"
	"os"
	"strconv"
)

func is_number(n string) bool {
	Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	if _, err := strconv.Atoi(n); err == nil {
		Info.Println("looks like it's a number..")
		return true
	}

	Info.Println("looks like it's not a number..")
	return false
}
