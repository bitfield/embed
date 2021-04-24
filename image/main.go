package main

import (
	_ "embed"
	"fmt"
	"net/http"
)

//go:embed rocket.png
var img []byte

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png")
		fmt.Fprintln(w, img)
	})
	http.ListenAndServe(":9999", nil)
}