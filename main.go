package main

import (
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
)

func main() {
	N := 100
	file, err := os.Create("housesOutputGo.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for i := 0; i < N; i++ {
		housesFile, err := os.Open("housesInput.csv")
		if err != nil {
			log.Fatal(err)
		}

		housesDF := dataframe.ReadCSV(housesFile)
		desc := housesDF.Describe()

		descString := desc.String()
		_, err = file.WriteString(descString + "\n")
		if err != nil {
			log.Fatal(err)
		}

		housesFile.Close()
	}
}
