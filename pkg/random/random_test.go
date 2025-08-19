package random

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestElement_IntArray(t *testing.T) {
	array := []int{1, 2, 3, 4, 5}
	element := Element(array)

	assert.True(t, slices.Contains(array, element))
}

func TestElement_StringArray(t *testing.T) {
	array := []string{"apple", "banana", "cherry"}
	element := Element(array)

	assert.True(t, slices.Contains(array, element))
}

func TestIndex_IntArray(t *testing.T) {
	array := []int{10, 20, 30, 40, 50}
	index := Index(array)

	assert.GreaterOrEqual(t, index, int64(0))
	assert.Less(t, index, int64(len(array)))
}

func TestIndex_StringArray(t *testing.T) {
	array := []string{"dog", "cat", "fish"}
	index := Index(array)

	assert.GreaterOrEqual(t, index, int64(0))
	assert.Less(t, index, int64(len(array)))
}

func TestElement_EmptyArray(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Ожидалась паника при передаче пустого массива")
		}
	}()

	_ = Element([]int{})
}

func TestIndex_EmptyArray(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Ожидалась паника при передаче пустого массива")
		}
	}()

	_ = Index([]int{})
}
