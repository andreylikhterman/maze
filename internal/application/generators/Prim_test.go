package generators

import (
	"testing"

	domain "maze/internal/domain"

	"github.com/stretchr/testify/assert"
)

func TestPrimGenerator_Initialize(t *testing.T) {
	generator := PrimGenerator{}
	height, width := 5, 5
	generator.maze = domain.NewMaze(height, width)

	generator.Initialize()

	for i := range height {
		for j := range width {
			coord := domain.NewCoordinate(i, j)
			assert.False(t, generator.isVisited[coord])
		}
	}

	assert.True(t, generator.heap.Empty())
}

func TestPrimGenerator_RandomWeights(t *testing.T) {
	generator := PrimGenerator{}
	height, width := 3, 3
	generator.maze = domain.NewMaze(height, width)

	weights := generator.randomWeights()

	expectedLength := height * width
	assert.Equal(t, expectedLength, len(weights))

	uniqueWeights := make(map[int]bool)
	for _, w := range weights {
		uniqueWeights[w] = true
	}

	assert.Equal(t, expectedLength, len(uniqueWeights))
}

func TestPrimGenerator_Neighbors(t *testing.T) {
	generator := PrimGenerator{}
	height, width := 3, 3
	generator.maze = domain.NewMaze(height, width)

	coord := domain.NewCoordinate(1, 1)
	neighbors := generator.neighbors(coord)

	expectedNeighbors := []domain.Coordinate{
		domain.NewCoordinate(0, 1),
		domain.NewCoordinate(1, 2),
		domain.NewCoordinate(2, 1),
		domain.NewCoordinate(1, 0),
	}

	assert.ElementsMatch(t, expectedNeighbors, neighbors)
}

func TestPrimGenerator_UnvisitedNeighbor(t *testing.T) {
	generator := PrimGenerator{}
	height, width := 3, 3
	generator.maze = domain.NewMaze(height, width)
	generator.Initialize()

	startCoord := domain.NewCoordinate(1, 1)
	visitedCoord := domain.NewCoordinate(0, 1)
	generator.isVisited[visitedCoord] = true

	neighbors := generator.unvisitedNeighbor(startCoord)

	expectedNeighbors := []domain.Coordinate{
		domain.NewCoordinate(1, 2),
		domain.NewCoordinate(2, 1),
		domain.NewCoordinate(1, 0),
	}

	assert.ElementsMatch(t, expectedNeighbors, neighbors)
}

func TestPrimGenerator_RandomCoordinate(t *testing.T) {
	generator := PrimGenerator{}
	height, width := 3, 3
	generator.maze = domain.NewMaze(height, width)

	coord := generator.randomCoordinate(generator.maze)

	assert.True(t, coord.Row() >= 0 && coord.Row() < height)
	assert.True(t, coord.Column() >= 0 && coord.Column() < width)
}
