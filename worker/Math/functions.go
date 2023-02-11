package Math

import (
	"math"
	"sync"
	"sync/atomic"
	"unsafe"
)

func AtomicAddFloat64(val *float64, delta float64) (new float64) {
	for {
		old := *val
		new = old + delta
		if atomic.CompareAndSwapUint64(
			(*uint64)(unsafe.Pointer(val)),
			math.Float64bits(old),
			math.Float64bits(new),
		) {
			break
		}
	}
	return
}

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func AtomicAdd[T Number](val *T, delta T) T {
	var p atomic.Pointer[T]
	var ans T
	for {
		p.Store(val)
		ans = *val + delta
		if p.CompareAndSwap(val, &ans) {
			break
		}
	}
	return ans

}

func DotProduct(seq1 *FloatSequence, seq2 *FloatSequence) (float64, error) {
	if seq1.GetLength() != seq2.GetLength() {
		return 0, ErrDifferentLength
	}

	var result float64 = 0

	for i := 0; i < seq1.GetLength(); i++ {
		a, err := seq1.GetItem(i)
		if err != nil {
			return 0, err
		}
		b, err := seq2.GetItem(i)
		if err != nil {
			return 0, err
		}

		result += a * b
	}

	return result, nil
}

func DotProductAsync(seq1 *FloatSequence, seq2 *FloatSequence) (float64, error) {
	if seq1.GetLength() != seq2.GetLength() {
		return 0, ErrDifferentLength
	}

	var result float64 = 0

	var wg sync.WaitGroup

	for i := 0; i < seq1.GetLength(); i++ {
		a, err := seq1.GetItem(i)
		if err != nil {
			return 0, err
		}
		b, err := seq2.GetItem(i)
		if err != nil {
			return 0, err
		}

		wg.Add(1)
		go func(a float64, b float64) {
			defer wg.Done()
			AtomicAdd(&result, a*b)
		}(a, b)
	}
	wg.Wait()

	return result, nil
}

func Min(a float64, args ...float64) float64 {
	min := a
	for _, val := range args {
		if val < min {
			min = val
		}
	}

	return min
}

func Max(a float64, args ...float64) float64 {
	max := a
	for _, val := range args {
		if val > max {
			max = val
		}
	}

	return max
}

func Sum(a float64, args ...float64) float64 {
	result := a
	for _, val := range args {
		result += val
	}

	return result
}

func Multiply(a float64, args ...float64) float64 {
	result := a
	for _, val := range args {
		result *= val
	}

	return result
}

func Divide(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrZeroDivision
	}

	return a / b, nil
}
