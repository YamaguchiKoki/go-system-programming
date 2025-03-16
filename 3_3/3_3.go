package main

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Create("output.zip")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := zip.NewWriter(file)
	defer writer.Close()

	content := strings.NewReader("Hello, World!")

	w, err := writer.Create("hello.txt")
	if err != nil {
		panic(err)
	}

	if _, err := io.Copy(w, content); err != nil {
		panic(err)
	}
}
