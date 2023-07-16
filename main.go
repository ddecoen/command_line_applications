package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	startTime := time.Now()
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
		log.Fatal(err)
	}

	// Extract housing cost values and calculate the average
	total := 0.0
	count := 0
	// Skip the first row (header)
	records = records[1:]
	for _, record := range records {
		// Extract the housing cost from the relevant column (e.g., index 0)
		housingCost, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Fatal(err)
		}

		total += housingCost
		count++
	}
	// Extract values from the desired column (e.g., index 2)
	columnIndex := 0
	var values []float64
	for _, record := range records {
		// Convert the field to a float64
		value, err := strconv.ParseFloat(record[columnIndex], 64)
		if err != nil {
			log.Fatal(err)
		}
		values = append(values, value)
	}
	// Find the maximum value
	max := values[0]
	for _, value := range values {
		if value > max {
			max = value
		}
	}

	// Extract housing income values and calculate the average
	totalinc := 0.0
	countinc := 0
	// Skip the first row (header)
	records = records[1:]
	for _, record := range records {
		housingIncome, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Fatal(err)
		}

		totalinc += housingIncome
		countinc++
	}
	// Extract housing age values and calculate the average
	totalage := 0.0
	countage := 0
	// Skip the first row (header)
	records = records[1:]
	for _, record := range records {
		housingAge, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			log.Fatal(err)
		}

		totalage += housingAge
		countage++
	}
	// Extract housing rooms values and calculate the average
	totalrooms := 0.0
	countrooms := 0
	// Skip the first row (header)
	records = records[1:]
	for _, record := range records {
		housingRooms, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			log.Fatal(err)
		}

		totalrooms += housingRooms
		countrooms++
	}
	// Extract housing bedrooms values and calculate the average
	totalbdr := 0.0
	countbdr := 0
	// Skip the first row (header)
	records = records[1:]
	for _, record := range records {
		housingBdr, err := strconv.ParseFloat(record[4], 64)
		if err != nil {
			log.Fatal(err)
		}

		totalbdr += housingBdr
		countbdr++
	}
	// Extract housing pop values and calculate the average
	totalpop := 0.0
	countpop := 0
	// Skip the first row (header)
	records = records[1:]
	for _, record := range records {
		housingPop, err := strconv.ParseFloat(record[5], 64)
		if err != nil {
			log.Fatal(err)
		}

		totalpop += housingPop
		countpop++
	}
	// Extract housing hh values and calculate the average
	totalhh := 0.0
	counthh := 0
	// Skip the first row (header)
	records = records[1:]
	for _, record := range records {
		housingHh, err := strconv.ParseFloat(record[6], 64)
		if err != nil {
			log.Fatal(err)
		}

		totalhh += housingHh
		counthh++
	}

	// Calculate the average housing metrics
	averageHousingCost := total / float64(count)
	averageHousingIncome := totalinc / float64(countinc)
	averageHousingAge := totalage / float64(countage)
	averageHousingRooms := totalrooms / float64(countrooms)
	averageHousingBdr := totalbdr / float64(countbdr)
	averageHousingPop := totalpop / float64(countpop)
	averageHousingHh := totalhh / float64(counthh)

	fmt.Printf("Average housing cost: $%.2f\n", averageHousingCost)
	fmt.Printf("Maximum housing value:, $%.2f\n", max)
	fmt.Printf("Average housing income: $%.2f\n", averageHousingIncome)
	fmt.Printf("Average housing age: %.2f\n", averageHousingAge)
	fmt.Printf("Average housing rooms: %.2f\n", averageHousingRooms)
	fmt.Printf("Average housing bedrooms: %.2f\n", averageHousingBdr)
	fmt.Printf("Average housing pop: %.2f\n", averageHousingPop)
	fmt.Printf("Average housing hh: %.2f\n", averageHousingHh)
	fmt.Println("Time to run summary statistics: ", time.Since(startTime))
}
