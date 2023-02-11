package Math

import (
	"testing"
)

var TestingFloatSequence1 = FloatSequence{
	length: 13,
	data:   []float64{14, 15, 17, 22, 99, 44, 29, 45, 51, 30, 87, 74, 3},
}

var TestingFloatSequence2 = FloatSequence{
	length: 13,
	data:   []float64{43, 31, 3, 74, 87, 30, 51, 45, 29, 44, 99, 22, 17},
}

const TestingDotProduct float64 = 29274
const TestingA float64 = 10
const TestingB float64 = 15
const TestingSum float64 = 25
const TestingProduct float64 = 150
const TestingDivision float64 = 1.5

func TestDotProduct(t *testing.T) {
	dp, err := DotProduct(&TestingFloatSequence1, &TestingFloatSequence2)

	if err != nil {
		t.Errorf("Got error %e", err)
	} else if dp != TestingDotProduct {
		t.Errorf("Expected %f, got %f", TestingDotProduct, dp)
	}

	_, err = DotProduct(&TestingFloatSequence1, &TestingFloatSequence)
	if err != ErrDifferentLength {
		t.Errorf("Expected ErrDifferentLength, got %e", err)
	}
}

func TestDotProductAsync(t *testing.T) {
	dp, err := DotProductAsync(&TestingFloatSequence1, &TestingFloatSequence2)

	if err != nil {
		t.Errorf("Got error %e", err)
	} else if dp != TestingDotProduct {
		t.Errorf("Expected %f, got %f", TestingDotProduct, dp)
	}

	_, err = DotProductAsync(&TestingFloatSequence1, &TestingFloatSequence)
	if err != ErrDifferentLength {
		t.Errorf("Expected ErrDifferentLength, got %e", err)
	}
}

func TestMin(t *testing.T) {
	min := Min(TestingA, TestingB)

	if min != TestingA {
		t.Errorf("Expected %f, got %f", TestingA, min)
	}
}

func TestMax(t *testing.T) {
	max := Max(TestingA, TestingB)

	if max != TestingB {
		t.Errorf("Expected %f, got %f", TestingB, max)
	}
}

func TestSum(t *testing.T) {
	sum := Sum(TestingA, TestingB)

	if sum != TestingSum {
		t.Errorf("Expected %f, got %f", TestingSum, sum)
	}
}

func TestProduct(t *testing.T) {
	pr := Multiply(TestingA, TestingB)

	if pr != TestingProduct {
		t.Errorf("Expected %f, got %f", TestingProduct, pr)
	}
}

func TestDivide(t *testing.T) {
	dv, err := Divide(TestingB, TestingA)

	if err != nil {
		t.Errorf("Got %e", err)
	} else if dv != TestingDivision {
		t.Errorf("Expected %f, got %f", TestingDivision, dv)
	}
}

func TestAtomicAdd(t *testing.T) {
	a := 5
	b := 4
	dv := AtomicAdd(&a, b)
	if a+b != dv {
		t.Errorf("Expected %d, got %d", a+b, dv)
	}
}
