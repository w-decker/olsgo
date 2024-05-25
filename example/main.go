package main

import (
	"log"

	"github.com/w-decker/olsgo"
)

func main() {

	// load data
	d, err := olsgo.LoadCSV("height_v_weight.csv")
	if err != nil {
		log.Fatalf("Error: %v \nCannot load file.")
	}

	// compute O-L-S
	reg := olsgo.OLS(d, "wt", "ht")

	// save
	reg.Save("output")

	// plot raw
	err = olsgo.PlotRaw("height_v_weight.csv", "plot.png", "wt", "ht")
	if err != nil {
		log.Fatalf("Error: %v \nCannot render plot.")
	}
}
