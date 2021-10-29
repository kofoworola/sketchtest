package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/draw", drawHandler)
	if err := http.ListenAndServe(":8085", nil); err != nil {
		log.Fatal(err)
	}
}
