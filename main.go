package main

import (
	"fizzbuzz/fizzbuzz"
	"net/http"
)

func main() {
	http.HandleFunc("/fizzbuzz", fizzbuzz.HandleFizzbuzz)
	http.ListenAndServe(":4000", nil)
}
