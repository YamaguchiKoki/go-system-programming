package main

import (
	"io"
	"os"
)

func main() {
	in, err := os.Open("old.txt")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	out, err := os.Create("new.txt")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	io.Copy(out, in)
}
