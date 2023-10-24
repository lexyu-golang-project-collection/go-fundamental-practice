package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	// Check if at least one CSV file path is provided as a command-line argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: " + os.Args[0] + " <csv_file1> [<csv_file2> ...]")
		return
	}

	// Create a slice to store all the data from the CSV files
	var allData []map[string]string

	// Iterate through the command-line arguments, starting from the second argument (index 1)
	for _, csvFile := range os.Args[1:] {
		// Open the CSV file
		file, err := os.Open(csvFile)
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

		// Append the data to the allData slice
		allData = append(allData, data...)
	}

	// Convert the combined data to JSON
	jsonData, err := json.MarshalIndent(allData, "", "	")
	if err != nil {
		log.Fatal(err)
	}

	// Write the JSON data to a single file
	err = os.WriteFile("combined_data.json", jsonData, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// Print a success message
	fmt.Println("Successfully converted and combined CSV files to JSON.")
}

func hardCodeVer() {
	// Open the CSV file
	file, err := os.Open("./data.csv")
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
	err = os.WriteFile("./data.json", jsonData, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// Print a success message
	fmt.Println("Successfully")
}
