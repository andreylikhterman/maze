package generators

import (
	"testing"

	domain "maze/internal/domain"
	algorithm "maze/pkg/algorithm"
	math "maze/pkg/math"

	"github.com/stretchr/testify/assert"
)

func TestKruskalGenerator_Edges(t *testing.T) {
	generator := KruskalGenerator{}
	height, width := 4, 4
	maze := domain.NewMaze(height, width)

	edges := generator.edges(maze, height, width)

	expectedEdges := 24
	assert.Equal(t, expectedEdges, len(edges))

	for _, edge := range edges {
		start := edge.Start().Coordinate()
		end := edge.End().Coordinate()
		assert.True(t, areNeighbors(start, end))
	}
}

func TestKruskalGenerator_DisjointSetOperations(t *testing.T) {
	disjointSet := algorithm.NewDisjointSet[domain.Coordinate]()

	coord1 := domain.NewCoordinate(0, 0)
	coord2 := domain.NewCoordinate(0, 1)
	coord3 := domain.NewCoordinate(1, 0)

	disjointSet.MakeSet(coord1)
	disjointSet.MakeSet(coord2)
	disjointSet.MakeSet(coord3)

	assert.NotEqual(t, disjointSet.Find(coord1), disjointSet.Find(coord2))

	disjointSet.Union(coord1, coord2)
	assert.Equal(t, disjointSet.Find(coord1), disjointSet.Find(coord2))

	assert.NotEqual(t, disjointSet.Find(coord1), disjointSet.Find(coord3))
}

func TestKruskalGenerator_ShuffleEdges(t *testing.T) {
	generator := KruskalGenerator{}
	height, width := 3, 3
	maze := domain.NewMaze(height, width)

	edges := generator.edges(maze, height, width)
	shuffledEdges := make([]domain.Edge, len(edges))
	copy(shuffledEdges, edges)
	shuffledEdges = algorithm.Shuffle(shuffledEdges)

	assert.NotEqual(t, edges, shuffledEdges)

	assert.Equal(t, len(edges), len(shuffledEdges))
}

func areNeighbors(a, b domain.Coordinate) bool {
	rowDiff := math.Abs(a.Row() - b.Row())
	colDiff := math.Abs(a.Column() - b.Column())

	return (rowDiff == 1 && colDiff == 0) || (rowDiff == 0 && colDiff == 1)
}
