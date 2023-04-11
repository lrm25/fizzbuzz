package main

import (
	"fizzbuzz/fizzbuzz"
	"os"
)

const ENV_SERVERPORT = "FIZZBUZZ_SERVERPORT"
const ENV_CLIENTURL = "FIZZBUZZ_CLIENTURL"

func main() {

	port := os.Getenv(ENV_SERVERPORT)
	if port == "" {
		panic("Port environment value not specified")
	}
	clientUrl := os.Getenv(ENV_CLIENTURL)
	if clientUrl == "" {
		panic("Client URL not specified")
	}
	server := fizzbuzz.NewFizzbuzzServer(port, clientUrl)
	server.Start()
}
