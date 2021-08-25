package problem

import "math"

func myPow(x float64, n int) float64 {
	pow := math.Pow(x, float64(n))
	return math.Round(pow*100000) / 100000
}
