package olsgo

// resourced from: https://www.geeksforgeeks.org/how-to-read-a-csv-file-in-golang/

import (
	"encoding/csv"
	"fmt"
	"os"
)

// read .csv file and returns a map where key-val pairs are header and observation data

func LoadCSV(filename string) (map[string][]string, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("could not open file %s: %v", filename, err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	headers, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("could not read headers: %v", err)
	}

	result := make(map[string][]string)
	for _, header := range headers {
		result[header] = []string{}
	}

	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, fmt.Errorf("could not read record: %v", err)
		}

		for i, value := range record {
			result[headers[i]] = append(result[headers[i]], value)
		}
	}

	return result, nil
}
