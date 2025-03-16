package main

import (
	"archive/zip"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename=output.zip")

	// レスポンスライターをzip.Writerでラップ
	writer := zip.NewWriter(w)
	defer writer.Close()

	// ZIPエントリの作成
	zipper, err := writer.Create("./3_4.go")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	src, err := os.Open("./3_4.go")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer src.Close()

	if _, err := io.Copy(zipper, src); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
