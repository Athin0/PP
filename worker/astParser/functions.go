package astParser

import "PP/worker/genericMath"

/*
Function substitutes
*/
var getMean = func(sequence *genericMath.FloatSequence) float64 {
	res, _ := sequence.GetMean()
	return res
}

var getVariance = func(sequence *genericMath.FloatSequence) float64 {
	res, _ := sequence.GetVariance()
	return res
}

var getStandardDeviation = func(sequence *genericMath.FloatSequence) float64 {
	res, _ := sequence.GetStandardDeviation()
	return res
}

var add = func(a, b float64) float64 {
	return genericMath.Sum(a, b)
}

var sub = func(a, b float64) float64 {
	return genericMath.Sum(a, -b)
}

var mul = func(a, b float64) float64 {
	return genericMath.Multiply(a, b)
}

var div = func(a, b float64) float64 {
	res, _ := genericMath.Divide(a, b)
	return res
}

var min = func(a, b float64) float64 {
	return genericMath.Min(a, b)
}

var max = func(a, b float64) float64 {
	return genericMath.Max(a, b)
}
