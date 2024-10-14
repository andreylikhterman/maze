package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbs_Int(t *testing.T) {
	assert.Equal(t, 5, Abs(-5))
	assert.Equal(t, 5, Abs(5))
	assert.Equal(t, 0, Abs(0))
}

func TestAbs_Int64(t *testing.T) {
	assert.Equal(t, int64(100), Abs(int64(-100)))
	assert.Equal(t, int64(100), Abs(int64(100)))
	assert.Equal(t, int64(0), Abs(int64(0)))
}

func TestAbs_Float32(t *testing.T) {
	assert.Equal(t, float32(3.14), Abs(float32(-3.14)))
	assert.Equal(t, float32(3.14), Abs(float32(3.14)))
	assert.Equal(t, float32(0), Abs(float32(0)))
}

func TestAbs_Float64(t *testing.T) {
	assert.Equal(t, 2.718, Abs(-2.718))
	assert.Equal(t, 2.718, Abs(2.718))
	assert.Equal(t, 0.0, Abs(0.0))
}

func TestAbs_BoundaryValues(t *testing.T) {
	assert.Equal(t, int64(9223372036854775807), Abs(int64(-9223372036854775807)))
	assert.Equal(t, int64(9223372036854775807), Abs(int64(9223372036854775807)))
}
