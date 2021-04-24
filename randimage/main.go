package main

import (
	"embed"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

//go:embed img/*
var images embed.FS

func main() {
	rand.Seed(time.Now().Unix())
	entries, err := images.ReadDir("img")
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL)
		imgPath := entries[rand.Intn(len(entries))]
		fPath := "img/" + imgPath.Name()
		fmt.Println("image", fPath)
		f, err := images.Open(fPath)
		if err != nil {
			http.Error(w, "no soup for you", http.StatusInternalServerError)
			return
		}
		defer f.Close()
		w.Header().Set("Content-Type", "image/png")
		io.Copy(w, f)
	})
	log.Fatal(http.ListenAndServe(":9999", nil))
}