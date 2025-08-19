package solvers

import (
	"testing"

	domain "maze/internal/domain"

	"github.com/stretchr/testify/assert"
)

func TestAStarSolver_Initialize(t *testing.T) {
	aStar := AStarSolver[int]{}
	height, width := 5, 5
	maze := domain.NewMaze(height, width)

	aStar.Initialize(maze)

	assert.True(t, aStar.heap.Empty())

	assert.Empty(t, aStar.parent)
	assert.Empty(t, aStar.distance)

	assert.Empty(t, aStar.closedSet)
}

func TestAStarSolver_Heuristic(t *testing.T) {
	aStar := AStarSolver[int]{}
	start := domain.NewCoordinate(0, 0)
	end := domain.NewCoordinate(3, 4)

	heuristic := aStar.heuristic(start, end)
	expectedHeuristic := 7
	assert.Equal(t, expectedHeuristic, heuristic)
}

func TestAStarSolver_Neighbors(t *testing.T) {
	aStar := AStarSolver[int]{}
	height, width := 3, 3
	maze := domain.NewMaze(height, width)

	aStar.Initialize(maze)

	coord := domain.NewCoordinate(1, 1)

	neighbors := aStar.Neighbors(coord)
	assert.Empty(t, neighbors)

	domain.DeleteWall(maze, coord, domain.NewCoordinate(1, 2))
	domain.DeleteWall(maze, coord, domain.NewCoordinate(2, 1))

	neighbors = aStar.Neighbors(coord)
	assert.Equal(t, 2, len(neighbors))
}

func TestAStarSolver_ReconstructPath(t *testing.T) {
	aStar := AStarSolver[int]{}
	height, width := 5, 5
	maze := domain.NewMaze(height, width)

	aStar.Initialize(maze)

	start := domain.NewCoordinate(0, 0)
	end := domain.NewCoordinate(4, 4)
	aStar.parent[end] = domain.NewCoordinate(3, 4)
	aStar.parent[domain.NewCoordinate(3, 4)] = domain.NewCoordinate(2, 4)
	aStar.parent[domain.NewCoordinate(2, 4)] = start

	path := aStar.reconstructPath(end)

	expectedPath := []domain.Coordinate{
		start,
		domain.NewCoordinate(2, 4),
		domain.NewCoordinate(3, 4),
		end,
	}

	assert.Equal(t, expectedPath, path)
}

func TestAStarSolver_IsVisited(t *testing.T) {
	aStar := AStarSolver[int]{}
	height, width := 3, 3
	maze := domain.NewMaze(height, width)

	aStar.Initialize(maze)

	coord := domain.NewCoordinate(1, 1)

	assert.False(t, aStar.isVisited(coord))

	aStar.closedSet[coord] = true
	assert.True(t, aStar.isVisited(coord))
}
