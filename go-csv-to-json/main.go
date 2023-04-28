package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	// Open the CSV file
	file, err := os.Open("data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read the CSV file
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Convert the records to a slice of maps
	var data []map[string]string
	headers := records[0]
	for _, record := range records[1:] {
		row := make(map[string]string)
		for i, value := range record {
			row[headers[i]] = value
		}
		data = append(data, row)
	}

	// Convert the data to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	// Write the JSON data to file
	err = os.WriteFile("data.json", jsonData, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// Print a success message
	fmt.Println("Successfully")
}
