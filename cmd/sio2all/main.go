package main

import (
	"fmt"
	"net/http"

	"github.com/creack/termios/raw"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi %s", r.URL.Path[1:])
}

func main() {
	raw.
		http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
