package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request ) {
	fmt.Fprintf(w, "hello!\n")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}