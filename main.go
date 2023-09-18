package main

import (
	"log"
	"net/http"
)

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Hello World!"))
}

func main() {
	http.HandleFunc("/", RequestHandler)

	log.Fatal(http.ListenAndServe(":9999", nil))
}
