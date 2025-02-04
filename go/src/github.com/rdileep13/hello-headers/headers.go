package main

import (
	"net/http"
	"fmt"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received Request at %v\n", time.Now())
	for k, v := range r.Header {
		fmt.Printf("%v: %v\n", k, v)
	}
	fmt.Printf("Done processing Request at %v\n", time.Now())
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9090", nil)
}
