package olsgo

// resourced from: https://www.geeksforgeeks.org/how-to-read-a-csv-file-in-golang/

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// read .csv file and returns a map where key-val pairs are header and observation data
func LoadCSV(filename string) (map[string][]float64, error) {

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

	result := make(map[string][]float64)
	for _, header := range headers {
		result[header] = []float64{}
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
			floatValue, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return nil, fmt.Errorf("could not parse value %s: %v", value, err)
			}
			result[headers[i]] = append(result[headers[i]], floatValue)
		}
	}

	return result, nil
}
