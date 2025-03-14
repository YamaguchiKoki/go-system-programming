package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")

	source := map[string]string{
		"Hello": "World",
	}
	zipper := gzip.NewWriter(w)
	defer zipper.Close()

	writer := io.MultiWriter(zipper, os.Stdout)

	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "  ")
	encoder.Encode(source)
}

func main() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
	log.Println("Server is running on port 8080")
}
