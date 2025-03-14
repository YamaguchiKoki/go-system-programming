package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

func main() {
	records := [][]string{
		{"first_name", "last_name", "age", "sex"},
		{"John", "Doe", "25", "male"},
		{"Jane", "Doe", "30", "female"},
	}

	f, err := os.Create("output.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	writer := io.MultiWriter(os.Stdout, f)

	w := csv.NewWriter(writer)

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatal(err)
		}
	}

	w.Flush()
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
