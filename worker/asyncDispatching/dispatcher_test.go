package asyncDispatching

import (
	"fmt"
	"testing"

	"PP/worker/Math"
)

var getMean = func(sequence *Math.FloatSequence) float64 {
	res, _ := sequence.GetMean()
	return res
}
var getDispersion = func(sequence *Math.FloatSequence) float64 {
	res, _ := sequence.GetVariance()
	return res
}
var getDispersionAsync = func(sequence *Math.FloatSequence) float64 {
	res, _ := sequence.GetStandardDeviation()
	return res
}

var TestingNode21 = Node{UnaryAction: getMean}
var TestingNode22 = Node{UnaryAction: getDispersion}
var TestingNode23 = Node{UnaryAction: getMean}
var TestingNode24 = Node{UnaryAction: getDispersionAsync}
var TestingNode11 = Node{
	Left:  &TestingNode21,
	Right: &TestingNode22,
	BinaryAction: func(a, b float64) float64 {
		return Math.Min(a, b)
	},
}
var TestingNode12 = Node{
	Left:  &TestingNode23,
	Right: &TestingNode24,
	BinaryAction: func(a, b float64) float64 {
		return Math.Max(a, b)
	},
}
var TestingRoot = Node{
	Left:  &TestingNode11,
	Right: &TestingNode12,
	BinaryAction: func(a, b float64) float64 {
		return Math.Multiply(a, b)
	},
}

const TestingResult = "0.0009774308"

func evalTraverse() float64 {
	TestingNode21.Parent = &TestingNode11
	TestingNode22.Parent = &TestingNode11
	TestingNode23.Parent = &TestingNode12
	TestingNode24.Parent = &TestingNode12

	TestingNode11.Parent = &TestingRoot
	TestingNode12.Parent = &TestingRoot

	disp := NewDispatcher(&TestingRoot)

	return Traverse(disp)
}

func TestTraverse(t *testing.T) {
	result := evalTraverse()

	if fmt.Sprintf("%.10f", result) != TestingResult {
		t.Errorf("Expected %s, got %s", TestingResult, fmt.Sprintf("%.10f", result))
	}
}

func BenchmarkTraverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		evalTraverse()
	}
}
