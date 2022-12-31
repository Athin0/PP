package asyncDispatching

import "errors"

var ErrNoItem = errors.New("no such item")
var ErrNoJobHash = errors.New("no job with such hash")
var ErrEmptyCache = errors.New("getting item from empty cache")
