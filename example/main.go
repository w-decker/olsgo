package main

import (
	"fmt"
	"log"

	"github.com/w-decker/olsgo"
)

func main() {

	// load data
	d, err := olsgo.LoadCSV("height_v_weight.csv")
	if err != nil {
		log.Fatalf("Error: %v \nCannot load file.")
	}

	// normalize
	d["wt"] = olsgo.ZScore(d["wt"])
	d["ht"] = olsgo.ZScore(d["ht"])

	// compute O-L-S
	reg := olsgo.OLS(d, "wt", "ht") // out := olsgo.OLS(data (output from olsgo.LoadCsv), x-variable, y-variable)

	// print output to console
	fmt.Println(reg.Format())

	// save
	reg.Save("output")

	// plot raw
	err = olsgo.PlotRaw("height_v_weight.csv", "plot3.png", "wt", "ht")
	if err != nil {
		log.Fatalf("Error: %v \nCannot render plot.")
	}
}
