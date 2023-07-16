package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	// Open the CSV file
	file, err := os.Open("../housesInput.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read the CSV records one by one
	for {
		// Read a record
		record, err := reader.Read()

		// Stop reading if we reach the end of the file
		if err == io.EOF {
			break
		}

		// Handle other errors
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// Process the record
		fmt.Println(record)
	}
}
