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
	for name, headers := range r.Header {
        	for _, h := range headers {
        		fmt.Fprintf(w, "%v: %v\n", name, h)
        	}
    	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9090", nil)
}
