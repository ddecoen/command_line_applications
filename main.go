package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"gonum.org/v1/gonum/stat"
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

	// Read all the records from the CSV file
	// Read a record
	records, err := reader.ReadAll()

	// Stop reading if we reach the end of the file
	if err != nil {
		log.Fatal(err)
	}

	//Extract values from the records
	var values []float64
	for _, record := range records {
		for _, field := range record {
			//Convert the field to float64
			value, err := strconv.ParseFloat(field, 64)
			if err != nil {
				log.Fatal(err)
			}
			values = append(values, value)
		}
	}

	// Process the record
	//fmt.Println(record)

	// Calculate summary statistics
	//min := math.Min(values)
	//max := math.Max(values)
	mean := stat.Mean(values, nil)
	median := stat.Quantile(0.5, stat.Empirical, values, nil)
	stddev := stat.StdDev(values, nil)
	variance := stat.Variance(values, nil)

	// Print the summary statistics
	//fmt.Println("Minimum:", min)
	//fmt.Println("Maximum:", max)
	fmt.Println("Mean:", mean)
	fmt.Println("Median:", median)
	fmt.Println("Standard Deviation:", stddev)
	fmt.Println("Variance:", variance)
}
