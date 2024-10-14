package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDisjointSet(t *testing.T) {
	dSet := NewDisjointSet[int]()

	assert.NotNil(t, dSet, "Expected NewDisjointSet to return a non-nil value")
	assert.Empty(t, dSet.parent, "New DisjointSet should have an empty parent map")
	assert.Empty(t, dSet.rank, "New DisjointSet should have an empty rank map")
}

func TestDisjointSet_MakeSet(t *testing.T) {
	dSet := NewDisjointSet[int]()

	dSet.MakeSet(1)
	assert.Equal(t, 1, dSet.Find(1))
	assert.Equal(t, 0, dSet.rank[1])
}

func TestDisjointSet_Union(t *testing.T) {
	dSet := NewDisjointSet[int]()

	dSet.MakeSet(1)
	dSet.MakeSet(2)

	dSet.Union(1, 2)

	assert.Equal(t, dSet.Find(1), dSet.Find(2))
}

func TestDisjointSet_Union_Rank(t *testing.T) {
	dSet := NewDisjointSet[int]()

	dSet.MakeSet(1)
	dSet.MakeSet(2)
	dSet.MakeSet(3)

	dSet.Union(1, 2)
	dSet.Union(2, 3)

	assert.Equal(t, dSet.Find(1), dSet.Find(3))
	assert.Equal(t, 1, dSet.rank[1])
	assert.Equal(t, 0, dSet.rank[2])
	assert.Equal(t, 0, dSet.rank[3])
}

func TestDisjointSet_Find_PathCompression(t *testing.T) {
	dSet := NewDisjointSet[int]()

	dSet.MakeSet(1)
	dSet.MakeSet(2)
	dSet.MakeSet(3)

	dSet.Union(1, 2)
	dSet.Union(2, 3)

	root := dSet.Find(3)

	assert.Equal(t, dSet.Find(2), root)
	assert.Equal(t, dSet.parent[2], root)
	assert.Equal(t, dSet.parent[3], root)
}
