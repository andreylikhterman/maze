package algorithm

import (
	"time"
)

type LCG struct {
	seed int64
}

func NewLCG() LCG {
	seed := time.Now().UnixNano()
	if seed < 0 {
		seed = -seed
	}

	return LCG{seed: seed}
}

func (lcg *LCG) Next() int64 {
	const (
		a = 1664525
		c = 1013904223
		m = (1 << 32) + 7
	)

	lcg.seed = ((a%m)*(lcg.seed%m) + c) % m

	return lcg.seed
}

func Shuffle[T any](array []T) []T {
	lcg := NewLCG()

	for i := len(array) - 1; i > 0; i-- {
		j := int(lcg.Next() % int64(i+1))
		array[i], array[j] = array[j], array[i]
	}

	return array
}
