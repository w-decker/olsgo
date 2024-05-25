package olsgo

import (
	"fmt"
	"image/color"
	"log"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

// plot raw data as scatter plot
func PlotRaw(p string, n string, x string, y string) error {

	d, err := LoadCSV(p)
	if err != nil {
		fmt.Errorf("could not read file: %v", err)
	}

	out, err := os.Create(n)
	if err != nil {
		log.Fatalf("Error: %v \nCould not create plot file.", err)
	}
	defer out.Close()

	plt := plot.New()

	pltdat := make(plotter.XYs, len(d[x]))

	for idx, n := range d[x] {
		pltdat[idx].X = n
	}
	for idx, n := range d[y] {
		pltdat[idx].Y = n
	}

	scatter, err := plotter.NewScatter(pltdat)
	if err != nil {
		fmt.Errorf("error: %v")
	}
	scatter.GlyphStyle.Shape = draw.CircleGlyph{}
	scatter.GlyphStyle.Color = color.RGBA{R: 102, B: 255, G: 102}
	scatter.GlyphStyle.Radius = vg.Points(4)
	plt.Add(scatter)
	plt.Title.Text = fmt.Sprintf("%s and %s", x, y)
	plt.X.Label.Text = x
	plt.Y.Label.Text = y

	wt, err := plt.WriterTo(300, 300, "png")
	if err != nil {
		log.Fatalf("Error: %v \nCould not create plot.", err)
	}

	_, err = wt.WriteTo(out)
	if err != nil {
		log.Fatalf("Error: %v \nCould not write plot.", err)
	}

	if err := out.Close(); err != nil {
		log.Fatalf("Error: %v \nCould not write plot.", err)
	}

	return err

}

// plot regression line on top of raw data
func PlotModel(o ols, p string, n string) error {

	d, err := LoadCSV(p)
	if err != nil {
		fmt.Errorf("could not read file: %v", err)
	}

	out, err := os.Create(n)
	if err != nil {
		log.Fatalf("Error: %v \nCould not create plot file.", err)
	}
	defer out.Close()

	plt := plot.New()

	pltdat := make(plotter.XYs, len(d[o.x]))

	for idx, n := range d[o.x] {
		pltdat[idx].X = n
	}
	for idx, n := range d[o.y] {
		pltdat[idx].Y = n
	}

	scatter, err := plotter.NewScatter(pltdat)
	if err != nil {
		fmt.Errorf("error: %v")
	}
	scatter.GlyphStyle.Shape = draw.CircleGlyph{}
	scatter.GlyphStyle.Color = color.RGBA{R: 102, B: 255, G: 102}
	scatter.GlyphStyle.Radius = vg.Points(4)
	plt.Add(scatter)
	plt.Title.Text = fmt.Sprintf("%s and %s", o.x, o.y)
	plt.X.Label.Text = o.x
	plt.Y.Label.Text = o.x

	// add regression line
	regressionLine := make(plotter.XYs, len(pltdat))
	for i, point := range pltdat {
		regressionLine[i].X = point.X
		regressionLine[i].Y = o.intercept + o.b1*point.X
	}

	line, err := plotter.NewLine(regressionLine)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}
	line.Color = color.RGBA{R: 0, B: 0, G: 0, A: 255}
	plt.Add(line)

	wt, err := plt.WriterTo(300, 300, "png")
	if err != nil {
		log.Fatalf("Error: %v \nCould not create plot.", err)
	}

	_, err = wt.WriteTo(out)
	if err != nil {
		log.Fatalf("Error: %v \nCould not write plot.", err)
	}

	if err := out.Close(); err != nil {
		log.Fatalf("Error: %v \nCould not write plot.", err)
	}

	return err

}
