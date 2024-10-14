package random

import "time"

func Element[T any](array []T) (value T) {
	if len(array) == 0 {
		panic("empty array")
	}

	index := (time.Now().UnixNano() / 1e3) % int64(len(array))

	return array[index]
}

func Index[T any](array []T) (index int64) {
	if len(array) == 0 {
		panic("empty array")
	}

	index = (time.Now().UnixNano() / 1e3) % int64(len(array))

	return index
}
