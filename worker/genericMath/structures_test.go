package genericMath

import (
	"fmt"
	"testing"
)

var TestingFloatSequence = FloatSequence{
	length: 15,
	data:   []float64{14, 15, 17, 22, 99, 44, 29, 45, 51, 30, 87, 74, 3, 31, 43},
}
var TestingEmptyFloatSequence = FloatSequence{
	length: 0,
	data:   []float64{},
}
var TestingSingleFloatSequence = FloatSequence{
	length: 1,
	data:   []float64{1},
}

const TestingLength = 15
const TestingMin float64 = 3
const TestingMax float64 = 99
const TestingSequenceSum float64 = 604
const TestingMean = "40.26667"
const TestingVariance = "774.35238"
const TestingStandardDeviation = "27.82719"
const TestingItem1 float64 = 14
const TestingItem7 float64 = 29
const TestingItem15 float64 = 43

func TestFloatSequence_GetLength(t *testing.T) {
	l := TestingFloatSequence.GetLength()

	if l != TestingLength {
		t.Errorf("Expected %d, got %d", TestingLength, l)
	}
}

func TestFloatSequence_GetMin(t *testing.T) {
	min, err := TestingFloatSequence.GetMin()

	if err != nil {
		t.Errorf("Got %e", err)
	} else if min != TestingMin {
		t.Errorf("Expected %f, got %f", TestingMin, min)
	}

	min, err = TestingSingleFloatSequence.GetMin()

	if err != nil {
		t.Errorf("Got %e", err)
	} else if min != TestingSingleFloatSequence.data[0] {
		t.Errorf("Expected %f, got %f", TestingSingleFloatSequence.data[0], min)
	}

	_, err = TestingEmptyFloatSequence.GetMin()

	if err != ErrEmptySequence {
		t.Errorf("Expected ErrEmptySequence, got %e", err)
	}
}

func TestFloatSequence_GetMax(t *testing.T) {
	max, err := TestingFloatSequence.GetMax()

	if err != nil {
		t.Errorf("Got %e", err)
	} else if max != TestingMax {
		t.Errorf("Expected %f, got %f", TestingMax, max)
	}

	max, err = TestingSingleFloatSequence.GetMax()

	if err != nil {
		t.Errorf("Got %e", err)
	} else if max != TestingSingleFloatSequence.data[0] {
		t.Errorf("Expected %f, got %f", TestingSingleFloatSequence.data[0], max)
	}

	_, err = TestingEmptyFloatSequence.GetMax()

	if err != ErrEmptySequence {
		t.Errorf("Expected ErrEmptySequence, got %e", err)
	}
}

func TestFloatSequence_Append(t *testing.T) {
	seq := FloatSequence{data: []float64{1, 1, 1, 1}}

	seq.Append(1)

	if len(seq.data) != 5 {
		t.Errorf("Expected len of 5, got len of %d", len(seq.data))
	} else if func() bool {
		for _, val := range seq.data {
			if val != 1 {
				return true
			}
		}
		return false
	}() {
		t.Errorf("Expected 1 in seq.data")
	}
}

func TestFloatSequence_GetItem(t *testing.T) {
	item1, err := TestingFloatSequence.GetItem(0)

	if err != nil {
		t.Errorf("Got %e", err)
	} else if item1 != TestingItem1 {
		t.Errorf("Expected %f, got %f", TestingItem1, item1)
	}

	item7, err := TestingFloatSequence.GetItem(6)

	if err != nil {
		t.Errorf("Got %e", err)
	} else if item7 != TestingItem7 {
		t.Errorf("Expected %f, got %f", TestingItem7, item7)
	}

	item15, err := TestingFloatSequence.GetItem(14)

	if err != nil {
		t.Errorf("Got %e", err)
	} else if item15 != TestingItem15 {
		t.Errorf("Expected %f, got %f", TestingItem15, item15)
	}

	_, err = TestingFloatSequence.GetItem(15)

	if err != ErrWrongIndex {
		t.Errorf("Expected ErrWrongIndex, got %e", err)
	}

	_, err = TestingFloatSequence.GetItem(-1)

	if err != ErrWrongIndex {
		t.Errorf("Expected ErrWrongIndex, got %e", err)
	}
}

func TestFloatSequence_GetSum(t *testing.T) {
	sum, err := TestingFloatSequence.GetSum()

	if err != nil {
		t.Errorf("Got %e", err)
	} else if sum != TestingSequenceSum {
		t.Errorf("Expected %f, got %f", TestingSequenceSum, sum)
	}

	_, err = TestingEmptyFloatSequence.GetSum()

	if err != ErrEmptySequence {
		t.Errorf("Expected ErrEmptySequence, got %e", err)
	}
}

func TestFloatSequence_GetMean(t *testing.T) {
	mean, err := TestingFloatSequence.GetMean()

	if err != nil {
		t.Errorf("Got %e", err)
	} else if fmt.Sprintf("%.5f", mean) != TestingMean {
		t.Errorf("Expected %s, got %s", TestingMean, fmt.Sprintf("%.5f", mean))
	}

	mean, err = TestingSingleFloatSequence.GetMean()

	if err != nil {
		t.Errorf("Got %e", err)
	} else if mean != TestingSingleFloatSequence.data[0] {
		t.Errorf("Expected %f, got %f", TestingSingleFloatSequence.data[0], mean)
	}

	_, err = TestingEmptyFloatSequence.GetMean()

	if err != ErrEmptySequence {
		t.Errorf("Expected ErrEmptySequence, got %e", err)
	}
}

func TestFloatSequence_GetVariance(t *testing.T) {
	v, err := TestingFloatSequence.GetVariance()

	if err != nil {
		t.Errorf("Got %e", err)
	} else if fmt.Sprintf("%.5f", v) != TestingVariance {
		t.Errorf("Expected %s, got %s", TestingVariance, fmt.Sprintf("%.5f", v))
	}

	v, err = TestingSingleFloatSequence.GetVariance()

	if err != nil {
		t.Errorf("Got %e", err)
	} else if v != 0 {
		t.Errorf("Expected %d, got %f", 0, v)
	}

	_, err = TestingEmptyFloatSequence.GetVariance()

	if err != ErrEmptySequence {
		t.Errorf("Expected ErrEmptySequence, got %e", err)
	}
}

func TestFloatSequence_GetVarianceAsync(t *testing.T) {
	v, err := TestingFloatSequence.GetVarianceAsync()

	if err != nil {
		t.Errorf("Got %e", err)
	} else if fmt.Sprintf("%.5f", v) != TestingVariance {
		t.Errorf("Expected %s, got %s", TestingVariance, fmt.Sprintf("%.5f", v))
	}

	v, err = TestingSingleFloatSequence.GetVarianceAsync()

	if err != nil {
		t.Errorf("Got %e", err)
	} else if v != 0.0 {
		t.Errorf("Expected %d, got %f", 0, v)
	}

	_, err = TestingEmptyFloatSequence.GetVarianceAsync()

	if err != ErrEmptySequence {
		t.Errorf("Expected ErrEmptySequence, got %e", err)
	}
}

func TestFloatSequence_GetStandardDeviation(t *testing.T) {
	v, err := TestingFloatSequence.GetStandardDeviation()

	if err != nil {
		t.Errorf("Got %e", err)
	} else if fmt.Sprintf("%.5f", v) != TestingStandardDeviation {
		t.Errorf("Expected %s, got %s", TestingStandardDeviation, fmt.Sprintf("%.5f", v))
	}

	v, err = TestingSingleFloatSequence.GetStandardDeviation()

	if err != nil {
		t.Errorf("Got %e", err)
	} else if v != 0 {
		t.Errorf("Expected %d, got %f", 0, v)
	}

	_, err = TestingEmptyFloatSequence.GetStandardDeviation()

	if err != ErrEmptySequence {
		t.Errorf("Expected ErrEmptySequence, got %e", err)
	}
}
