package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello World!")
}

func Greet(w http.ResponseWriter, r *http.Request) {
	query, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Bad Request")
		return
	}

	name := query.Get("name")
	if len(name) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Missing name parameter")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello %s", name)
}

func V2Greet(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "V2 Hello World!")
}

func V3Greet(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "V3 Hello World!")
}

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/greet", Greet)
	http.HandleFunc("/v2greet", V2Greet)
	http.HandleFunc("/v3greet", V3Greet)

	log.Fatal(http.ListenAndServe(":9999", nil))
}
