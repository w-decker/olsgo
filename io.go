package olsgo

// resourced from: https://www.geeksforgeeks.org/how-to-read-a-csv-file-in-golang/

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func (o *ols) Format() string {
	fs := "OLSGO Output \n"
	fs += fmt.Sprintf("Intercept: %-15.4f", o.intercept)
	fs += fmt.Sprintf("B1: %-15.4f", o.b1)
	fs += fmt.Sprintf("Pearson's r: %-15.4f", o.r)
	fs += fmt.Sprintf("Residual variance: %-15.4f", o.R2)

	return fs
}

func (o *ols) Save(n string) {
	data := []byte(o.Format())
	err := os.WriteFile(n+".txt", data, 0666)

	if err != nil {
		panic(err)
	}
	fmt.Println("Output saved.")
}

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
