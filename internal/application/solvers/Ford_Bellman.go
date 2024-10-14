package solvers

import (
	domain "MazeApp/internal/domain"
)

const (
	INF = 10000
)

type FordBellmanSolver[T numbers] struct {
	maze     *domain.Maze
	parent   map[domain.Coordinate]domain.Coordinate
	distance map[domain.Coordinate]T
	edges    []domain.Edge
}

func (fordBellman *FordBellmanSolver[T]) reconstructPath(end domain.Coordinate) []domain.Coordinate {
	path := make([]domain.Coordinate, 0)
	path = append(path, end)
	current, ok := fordBellman.parent[end]

	for ok {
		path = append(path, current)
		current, ok = fordBellman.parent[current]
	}

	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}

func (fordBellman *FordBellmanSolver[T]) Solve(maze domain.Maze, start, end domain.Coordinate) []domain.Coordinate {
	fordBellman.maze = &maze
	fordBellman.initialize()
	fordBellman.distance[start] = 0

	for i := 0; i < maze.Height()*maze.Width()-1; i++ {
		for _, edge := range fordBellman.edges {
			start := edge.Start().Coordinate()
			end := edge.End().Coordinate()

			if fordBellman.distance[end] > fordBellman.distance[start]+T(edge.End().Surface()) {
				fordBellman.distance[end] = fordBellman.distance[start] + T(edge.End().Surface())
				fordBellman.parent[end] = start
			}
		}
	}

	if fordBellman.distance[end] != INF {
		return fordBellman.reconstructPath(end)
	}

	return []domain.Coordinate{}
}

func (fordBellman *FordBellmanSolver[T]) initialize() {
	fordBellman.parent = make(map[domain.Coordinate]domain.Coordinate)
	fordBellman.distance = make(map[domain.Coordinate]T)
	fordBellman.edges = fordBellman.Edges()

	for i := 0; i < fordBellman.maze.Height(); i++ {
		for j := 0; j < fordBellman.maze.Width(); j++ {
			fordBellman.distance[domain.NewCoordinate(i, j)] = INF
		}
	}
}

func (fordBellman FordBellmanSolver[T]) Edges() []domain.Edge {
	edges := []domain.Edge{}

	for i := 0; i < fordBellman.maze.Height(); i++ {
		for j := 0; j < fordBellman.maze.Width(); j++ {
			for _, neighbor := range fordBellman.Neighbors(domain.NewCoordinate(i, j)) {
				edges = append(edges, domain.NewEdge(fordBellman.maze, domain.NewCoordinate(i, j), neighbor.Coordinate()))
			}
		}
	}

	return edges
}

func (fordBellman *FordBellmanSolver[T]) inBounds(coordinate domain.Coordinate) bool {
	return coordinate.Column() >= 0 && coordinate.Column() < fordBellman.maze.Width() &&
		coordinate.Row() >= 0 && coordinate.Row() < fordBellman.maze.Height()
}

func (fordBellman *FordBellmanSolver[T]) Neighbors(coordinate domain.Coordinate) []*domain.Cell {
	var neighbors []*domain.Cell

	distance := []domain.Coordinate{
		domain.NewCoordinate(-1, 0),
		domain.NewCoordinate(0, 1),
		domain.NewCoordinate(1, 0),
		domain.NewCoordinate(0, -1),
	}

	current := fordBellman.maze.Cell(coordinate)

	for i, dist := range distance {
		NeighborCoordinate := domain.Sum(coordinate, dist)
		if fordBellman.inBounds(NeighborCoordinate) && !current.Walls()[i] {
			neighbors = append(neighbors, fordBellman.maze.Cell(NeighborCoordinate))
		}
	}

	return neighbors
}
