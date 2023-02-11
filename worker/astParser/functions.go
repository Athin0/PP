package astParser

import "PP/worker/Math"

/*
Function substitutes
*/
var getMean = func(sequence *Math.FloatSequence) float64 {
	res, _ := sequence.GetMean()
	return res
}

var getVariance = func(sequence *Math.FloatSequence) float64 {
	res, _ := sequence.GetVariance()
	return res
}

var getStandardDeviation = func(sequence *Math.FloatSequence) float64 {
	res, _ := sequence.GetStandardDeviation()
	return res
}

var add = func(a, b float64) float64 {
	return Math.Sum(a, b)
}

var sub = func(a, b float64) float64 {
	return Math.Sum(a, -b)
}

var mul = func(a, b float64) float64 {
	return Math.Multiply(a, b)
}

var div = func(a, b float64) float64 {
	res, _ := Math.Divide(a, b)
	return res
}

var min = func(a, b float64) float64 {
	return Math.Min(a, b)
}

var max = func(a, b float64) float64 {
	return Math.Max(a, b)
}
