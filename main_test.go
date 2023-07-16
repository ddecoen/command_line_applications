package main

import (
	"encoding/csv"
	"log"
	"os"
	"testing"
)

func TestReadCSV(t *testing.T) {
	// Open the CSV file
	file, err := os.Open("../housesInput.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read all records from the CSV file
	records, err := reader.ReadAll()
	if err != nil {
		t.Errorf("Failed to read CSV file: %s", err)
	}

	// Check if records are not empty
	if len(records) == 0 {
		t.Error("No records found in CSV file")
	}

	// Check the number of columns in each record - should be 7
	expectedColumns := 7
	for _, record := range records {
		if len(record) != expectedColumns {
			t.Errorf("Unexpected number of columns. Expected: %d, Got: %d", expectedColumns, len(record))
		}
	}
}
