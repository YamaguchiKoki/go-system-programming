package main

import (
	"crypto/rand"
	"os"
)

func main() {
	buffer := make([]byte, 1024)
	_, err := rand.Read(buffer)
	if err != nil {
		panic(err)
	}

	file, err := os.Create("rand.bin")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if _, err := file.Write(buffer); err != nil {
		panic(err)
	}
}
