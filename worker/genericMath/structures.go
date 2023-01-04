package genericMath

import (
	"math"
	"sync"
)

type FloatSequence struct {
	length int
	data   []float64
}

func (seq *FloatSequence) GetLength() int {
	return seq.length
}
func NewFloatSequence(data []float64) *FloatSequence {
	return &FloatSequence{data: data, length: len(data)}
}

func (seq *FloatSequence) GetMin() (float64, error) {
	if seq.length == 0 {
		return 0, ErrEmptySequence
	}

	if seq.length == 1 {
		return seq.data[0], nil
	}

	return Min(seq.data[0], seq.data[1:]...), nil
}

func (seq *FloatSequence) GetMax() (float64, error) {
	if seq.length == 0 {
		return 0, ErrEmptySequence
	}

	if seq.length == 1 {
		return seq.data[0], nil
	}

	return Max(seq.data[0], seq.data[1:]...), nil
}

func (seq *FloatSequence) Append(item float64) {
	seq.data = append(seq.data, item)
	seq.length++
}

func (seq *FloatSequence) GetItem(index int) (float64, error) {
	if index < 0 || index >= seq.length {
		return 0, ErrWrongIndex
	}

	return seq.data[index], nil
}

func (seq *FloatSequence) GetSum() (float64, error) {
	if seq.length == 0 {
		return 0, ErrEmptySequence
	}

	if seq.length == 1 {
		return seq.data[0], nil
	}

	result := 0.0
	for _, val := range seq.data {
		result += val
	}

	return result, nil
}

func (seq *FloatSequence) GetMean() (float64, error) {
	sum, err := seq.GetSum()
	if err != nil {
		return 0, err
	}

	return sum / float64(seq.length), nil
}

func (seq *FloatSequence) GetVarianceAsync() (float64, error) {
	if seq.length == 0 {
		return 0, ErrEmptySequence
	}

	if seq.length == 1 {
		return 0, nil
	}

	result := 0.0
	mean, err := seq.GetMean()

	if err != nil {
		return 0, err
	}

	var wg sync.WaitGroup

	for _, val := range seq.data {
		wg.Add(1)

		go func(value float64) {
			defer wg.Done()
			AtomicAdd(&result, math.Pow(value-mean, 2))
		}(val)
	}

	wg.Wait()

	return result / float64(seq.length-1), nil
}

func (seq *FloatSequence) GetVariance() (float64, error) {
	if seq.length == 0 {
		return 0, ErrEmptySequence
	}

	if seq.length == 1 {
		return 0, nil
	}

	result := 0.0
	mean, err := seq.GetMean()

	if err != nil {
		return 0, err
	}

	for _, val := range seq.data {
		result += math.Pow(val-mean, 2)
	}

	return result / float64(seq.length-1), nil
}

func (seq *FloatSequence) GetStandardDeviation() (float64, error) {
	variance, err := seq.GetVariance()
	if err != nil {
		return 0, err
	}

	return math.Sqrt(variance), nil
}
