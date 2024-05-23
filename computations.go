package olsgo

import "math"

// sum over 1d slice
func sum(x []float64) float64 {
	sum := 0.0
	for i := 0; i < len(x); i++ {
		sum += float64(x[i])
	}
	return sum
}

// compute mean over 1d slice
func mean(x []float64) float64 {
	n := float64(len(x))
	s := sum(x)
	avg := s / n

	return avg
}

// compute total sums of squares for 1d slice
func tss(x []float64) float64 {
	var t []float64
	mu := mean(x)
	for _, v := range x {
		d := v - mu
		t = append(t, d*d)
	}

	ss := sum(t)
	return ss
}

// compute std over 1d slice
func std(x []float64) float64 {
	numer := tss(x)
	n := float64(len(x))

	sigma := math.Sqrt((numer / n))

	return sigma
}

func pearsonr(x []float64, y []float64) float64 {

	n := float64(len(x))
	m := make([]float64, len(x))
	for idx, n := range x {
		m = append(m, n*float64(y[idx]))
	}

	numer := sum(m)

	return numer/n - 1
}
