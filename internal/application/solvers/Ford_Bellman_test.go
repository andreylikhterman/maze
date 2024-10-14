package solvers

import (
	"testing"

	domain "MazeApp/internal/domain"

	"github.com/stretchr/testify/assert"
)

func TestFordBellmanSolver_Initialize(t *testing.T) {
	fordBellman := FordBellmanSolver[int]{}
	height, width := 5, 5

	fordBellman.maze = domain.NewMaze(height, width)
	fordBellman.initialize()

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			coord := domain.NewCoordinate(i, j)
			assert.Equal(t, INF, fordBellman.distance[coord])
		}
	}

	assert.Empty(t, fordBellman.parent)
}

func TestFordBellmanSolver_ReconstructPath(t *testing.T) {
	fordBellman := FordBellmanSolver[int]{}
	height, width := 5, 5

	fordBellman.maze = domain.NewMaze(height, width)
	fordBellman.initialize()

	start := domain.NewCoordinate(0, 0)
	end := domain.NewCoordinate(4, 4)
	fordBellman.parent[end] = domain.NewCoordinate(3, 4)
	fordBellman.parent[domain.NewCoordinate(3, 4)] = domain.NewCoordinate(2, 4)
	fordBellman.parent[domain.NewCoordinate(2, 4)] = start

	path := fordBellman.reconstructPath(end)

	expectedPath := []domain.Coordinate{
		start,
		domain.NewCoordinate(2, 4),
		domain.NewCoordinate(3, 4),
		end,
	}

	assert.Equal(t, expectedPath, path)
}

func TestFordBellmanSolver_Edges(t *testing.T) {
	fordBellman := FordBellmanSolver[int]{}
	height, width := 3, 3
	fordBellman.maze = domain.NewMaze(height, width)

	center := domain.NewCoordinate(1, 1)
	domain.DeleteWall(fordBellman.maze, center, domain.NewCoordinate(1, 2))
	domain.DeleteWall(fordBellman.maze, center, domain.NewCoordinate(2, 1))
	domain.DeleteWall(fordBellman.maze, center, domain.NewCoordinate(0, 1))
	domain.DeleteWall(fordBellman.maze, center, domain.NewCoordinate(1, 0))

	fordBellman.initialize()

	edges := fordBellman.Edges()

	var centerEdges []domain.Edge

	for _, edge := range edges {
		if edge.Start().Coordinate() == center {
			centerEdges = append(centerEdges, edge)
		}
	}

	assert.Equal(t, 4, len(centerEdges))
}

func TestFordBellmanSolver_NoSolution(t *testing.T) {
	fordBellman := FordBellmanSolver[int]{}
	height, width := 3, 3
	maze := domain.NewMaze(height, width)

	start := domain.NewCoordinate(0, 0)
	end := domain.NewCoordinate(2, 2)

	path := fordBellman.Solve(*maze, start, end)
	assert.Empty(t, path)
}
