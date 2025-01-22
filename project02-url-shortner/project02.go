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
	urlsMap := map[string]string{
		"gophercises":   "https://gophercises.com",
		"codeasthetics": "https://www.youtube.com/@CodeAesthetic",
	}

	fmt.Println("Route: ", r.URL.Path[1:])
	path := r.URL.Path[1:]
	longUrl, exists := urlsMap[path]

	if exists {
		http.Redirect(w, r, longUrl, http.StatusFound)

	} else {
		http.NotFound(w, r)
	}
}
