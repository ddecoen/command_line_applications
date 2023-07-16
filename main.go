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

	// Find the minimum value
	min := values[0]
	for _, value := range values {
		if value < min {
			min = value
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
	// Extract values from the desired column (e.g., housing income)
	columnIndexinc := 1
	var valuesinc []float64
	for _, record := range records {
		// Convert the field to a float64
		valueinc, err := strconv.ParseFloat(record[columnIndexinc], 64)
		if err != nil {
			log.Fatal(err)
		}
		valuesinc = append(valuesinc, valueinc)
	}
	// Find the maximum value
	maxinc := valuesinc[0]
	for _, valueinc := range valuesinc {
		if valueinc > maxinc {
			maxinc = valueinc
		}
	}

	// Find the minimum value
	mininc := valuesinc[0]
	for _, valueinc := range valuesinc {
		if valueinc < mininc {
			mininc = valueinc
		}
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
	// Extract values from the desired column (e.g., housing age)
	columnIndexage := 2
	var valuesage []float64
	for _, record := range records {
		// Convert the field to a float64
		valueage, err := strconv.ParseFloat(record[columnIndexage], 64)
		if err != nil {
			log.Fatal(err)
		}
		valuesage = append(valuesage, valueage)
	}
	// Find the maximum value
	maxage := valuesage[0]
	for _, valueage := range valuesage {
		if valueage > maxage {
			maxage = valueage
		}
	}
	// Find the minimum value
	minage := valuesage[0]
	for _, valueage := range valuesage {
		if valueage < minage {
			minage = valueage
		}
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
	// Extract values from the desired column (e.g., housing rooms)
	columnIndexrooms := 3
	var valuesrooms []float64
	for _, record := range records {
		// Convert the field to a float64
		valuerooms, err := strconv.ParseFloat(record[columnIndexrooms], 64)
		if err != nil {
			log.Fatal(err)
		}
		valuesrooms = append(valuesrooms, valuerooms)
	}
	// Find the maximum value
	maxrooms := valuesrooms[0]
	for _, valuerooms := range valuesrooms {
		if valuerooms > maxrooms {
			maxrooms = valuerooms
		}
	}
	// Find the minimum value
	minrooms := valuesrooms[0]
	for _, valuerooms := range valuesrooms {
		if valuerooms < minrooms {
			minrooms = valuerooms
		}
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
	// Extract values from the desired column (e.g., housing bedrooms)
	columnIndexbdr := 4
	var valuesbdr []float64
	for _, record := range records {
		// Convert the field to a float64
		valuebdr, err := strconv.ParseFloat(record[columnIndexbdr], 64)
		if err != nil {
			log.Fatal(err)
		}
		valuesbdr = append(valuesbdr, valuebdr)
	}
	// Find the maximum value
	maxbdr := valuesbdr[0]
	for _, valuebdr := range valuesbdr {
		if valuebdr > maxbdr {
			maxbdr = valuebdr
		}
	}
	// Find the minimum value
	minbdr := valuesbdr[0]
	for _, valuebdr := range valuesbdr {
		if valuebdr < minbdr {
			minbdr = valuebdr
		}
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
	// Extract values from the desired column (e.g., housing pop)
	columnIndexpop := 5
	var valuespop []float64
	for _, record := range records {
		// Convert the field to a float64
		valuepop, err := strconv.ParseFloat(record[columnIndexpop], 64)
		if err != nil {
			log.Fatal(err)
		}
		valuespop = append(valuespop, valuepop)
	}
	// Find the maximum value
	maxpop := valuespop[0]
	for _, valuepop := range valuespop {
		if valuepop > maxpop {
			maxpop = valuepop
		}
	}
	// Find the minimum value
	minpop := valuespop[0]
	for _, valuepop := range valuespop {
		if valuepop < minpop {
			minpop = valuepop
		}
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
	fmt.Printf("Maximum housing value: $%.2f\n", max)
	fmt.Printf("Minimum housing value: $%.2f\n", min)
	fmt.Printf("Average housing income: $%.2f\n", averageHousingIncome)
	fmt.Printf("Maximum housing income: $%.2f\n", maxinc)
	fmt.Printf("Minimum housing income: $%.2f\n", mininc)
	fmt.Printf("Average housing age: %.2f\n", averageHousingAge)
	fmt.Printf("Maximum housing age: %.2f\n", maxage)
	fmt.Printf("Minimum housing age: %.2f\n", minage)
	fmt.Printf("Average housing rooms: %.2f\n", averageHousingRooms)
	fmt.Printf("Maximum housing rooms: %.2f\n", maxrooms)
	fmt.Printf("Minimum housing rooms: %.2f\n", minrooms)
	fmt.Printf("Average housing bedrooms: %.2f\n", averageHousingBdr)
	fmt.Printf("Maximum housing bedrooms: %.2f\n", maxbdr)
	fmt.Printf("Minimum housing bedrooms: %.2f\n", minbdr)
	fmt.Printf("Average housing pop: %.2f\n", averageHousingPop)
	fmt.Printf("Maximum housing pop: %.2f\n", maxpop)
	fmt.Printf("Minimum housing pop: %.2f\n", minpop)
	fmt.Printf("Average housing hh: %.2f\n", averageHousingHh)
	fmt.Println("Time to run summary statistics: ", time.Since(startTime))
}
