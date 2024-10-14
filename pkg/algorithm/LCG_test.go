package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLCG_SeedIsNonNegative(t *testing.T) {
	lcg := NewLCG()

	assert.GreaterOrEqual(t, lcg.seed, int64(0))
}

func TestNext_ReturnsInt64(t *testing.T) {
	lcg := NewLCG()

	result := lcg.Next()

	assert.GreaterOrEqual(t, result, int64(0))
}

func TestNext_ProducesDifferentValues(t *testing.T) {
	lcg := NewLCG()
	first := lcg.Next()
	second := lcg.Next()

	assert.NotEqual(t, first, second)
}

func TestShuffle_ProducesDifferentOrder(t *testing.T) {
	array := []int{1, 2, 3, 4, 5}
	original := make([]int, len(array))
	copy(original, array)

	shuffled := Shuffle(array)

	assert.NotEqual(t, original, shuffled)
	assert.Len(t, original, len(shuffled))
}

func TestShuffle_AllElementsArePresent(t *testing.T) {
	array := []int{1, 2, 3, 4, 5}
	shuffled := Shuffle(array)

	elementMap := make(map[int]bool)
	for _, v := range shuffled {
		elementMap[v] = true
	}

	for _, v := range array {
		assert.True(t, elementMap[v])
	}
}
