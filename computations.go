package olsgo

import "math"

// output
type ols struct {
	intercept float64
	b1        float64
	r         float64
	R2        float64
}

// sum over 1d slice
func Sum(x []float64) float64 {
	var sum float64
	for _, v := range x {
		sum += v
	}
	return sum
}

// compute mean over 1d slice
func Mean(x []float64) float64 {
	return Sum(x) / float64(len(x))
}

// Standardize 1d slize
func ZScore(x []float64) []float64 {
	mu := Mean(x)
	sigma := Std(x)
	z := make([]float64, len(x))
	for i, v := range x {
		z[i] = (v - mu) / sigma
	}
	return z
}

// compute total sums of squares for 1d slice
func Tss(y []float64) float64 {
	var t []float64
	mu := Mean(y)
	for _, v := range y {
		d := v - mu
		t = append(t, d*d)
	}

	sst := Sum(t)
	return sst
}

// compute std over 1d slice
func Std(x []float64) float64 {
	numer := Tss(x)
	n := float64(len(x))
	return math.Sqrt(numer / (n - 1))
}

// Compute Pearson's r for two 1d slices.
func PearsonR(x []float64, y []float64) float64 {
	if len(x) != len(y) {
		return 0 // or handle error appropriately
	}

	xmu := Mean(x)
	ymu := Mean(y)

	var numer, xdenom, ydenom float64

	for i := range x {
		numer += (x[i] - xmu) * (y[i] - ymu)
		xdenom += (x[i] - xmu) * (x[i] - xmu)
		ydenom += (y[i] - ymu) * (y[i] - ymu)
	}

	r := numer / math.Sqrt(xdenom*ydenom)

	return r
}

func B1(x []float64, y []float64) float64 {
	r := PearsonR(x, y)
	sx := Std(x)
	sy := Std(y)

	b := r * (sy / sx)

	return b
}

// Calculate b0
func Intercept(x []float64, y []float64) float64 {
	b1 := B1(x, y)
	xbar := Mean(x)
	ybar := Mean(y)

	i := ybar - (b1 * xbar)

	return i
}

// Compute y-hat
func YHat(x []float64, y []float64) []float64 {
	inter := Intercept(x, y)
	b1 := B1(x, y)
	yhat := make([]float64, len(x))
	for i := range x {
		yhat[i] = inter + b1*x[i]
	}
	return yhat
}

// compute sums of squares explained for 1d slice
func Sse(y []float64, yhat []float64) float64 {
	var sse float64
	for i := range y {
		d := y[i] - yhat[i]
		sse += d * d
	}
	return sse
}

// calculate R2 from raw data
func R2(x []float64, y []float64) float64 {
	sst := Tss(y)
	yhat := YHat(x, y)
	sse := Sse(y, yhat)

	r2 := sse / sst
	return r2
}

// Compute ordinary-least-squares regression over x and y variable
func OLS(d map[string][]float64, x string, y string) ols {
	xv := d[x]
	yv := d[y]

	reg := ols{
		intercept: Intercept(xv, yv),
		b1:        B1(xv, yv),
		r:         PearsonR(xv, yv),
		R2:        R2(xv, yv),
	}

	return reg

}
