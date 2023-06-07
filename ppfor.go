package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	// Register pprof endpoints
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	// Your program logic here
	// ...
}

