package Math

import (
	"errors"
)

var ErrWrongIndex = errors.New("unable to get item by index")
var ErrDifferentLength = errors.New("sequences must be of equal length")
var ErrEmptySequence = errors.New("sequence shouldn't be empty to get this data")
var ErrZeroDivision = errors.New("cannot divide by zero")
var ErrUninitializedData = errors.New("uninitialized data in sequence")
