package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("servers running on port 8080")
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/advance" {
		fmt.Fprintln(w, "🥷")
	} else {
		fmt.Fprintln(w, r.URL)
	}

	urlsMap := map[string]string{
		"gophercises": "https://gophercises.com",
		"shazTheDev":  "https://shazTheDev.com",
	}

}
