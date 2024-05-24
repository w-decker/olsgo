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
	sum := 0.0
	for i := 0; i < len(x); i++ {
		sum += float64(x[i])
	}
	return sum
}

// compute mean over 1d slice
func Mean(x []float64) float64 {
	n := float64(len(x))
	s := Sum(x)
	avg := s / n

	return avg
}

// Standardize 1d slize
func ZScore(x []float64) []float64 {
	var z []float64
	mu := Mean(x)
	sigma := Std(x)
	for _, n := range x {
		z = append(z, ((n - mu) / sigma))
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

	sigma := math.Sqrt((numer / n))

	return sigma
}

// Compute Pearson's r for two 1d slices.
func PearsonR(x []float64, y []float64) float64 {

	n := float64(len(x))
	m := make([]float64, len(x))
	for idx, n := range x {
		m = append(m, n*float64(y[idx]))
	}

	numer := Sum(m)

	return numer/n - 1
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
	var t []float64
	ybar := Mean(y)
	for i := range y {
		d := yhat[i] - ybar
		t = append(t, d*d)
	}
	sse := Sum(t)
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
// func OLS(d map[string][]float64, x string, y string)
