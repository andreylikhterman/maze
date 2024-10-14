package solvers

import (
	domain "MazeApp/internal/domain"
	algorithm "MazeApp/pkg/algorithm"
	math "MazeApp/pkg/math"
)

type numbers interface {
	int | int16 | int32 | int64 | float32 | float64
}

type AStarSolver[T numbers] struct {
	maze      *domain.Maze
	heap      algorithm.Heap[domain.Coordinate, T]
	parent    map[domain.Coordinate]domain.Coordinate
	distance  map[domain.Coordinate]T
	closedSet map[domain.Coordinate]bool
}

func (aStar *AStarSolver[T]) heuristic(start, end domain.Coordinate) int {
	return math.Abs(start.Row()-end.Row()) + math.Abs(start.Column()-end.Column())
}

func (aStar *AStarSolver[T]) reconstructPath(end domain.Coordinate) []domain.Coordinate {
	path := make([]domain.Coordinate, 0)
	path = append(path, end)

	current, ok := aStar.parent[end]
	for ok {
		path = append(path, current)
		current, ok = aStar.parent[current]
	}

	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}

func (aStar *AStarSolver[T]) Solve(maze domain.Maze, start, end domain.Coordinate) []domain.Coordinate {
	aStar.Initialize(&maze)

	aStar.heap.Insert(start, T(aStar.heuristic(start, end)))
	aStar.distance[start] = 0

	for !aStar.heap.Empty() {
		current, _ := aStar.heap.Min()

		if current == end {
			return aStar.reconstructPath(end)
		}

		aStar.heap.ExtractMin()
		aStar.closedSet[current] = true

		for _, neighbor := range aStar.Neighbors(current) {
			neighborDist := aStar.distance[current] + T(neighbor.Surface())
			neighborCoord := neighbor.Coordinate()

			if aStar.isVisited(neighborCoord) {
				continue
			}

			if !aStar.isVisited(neighborCoord) || neighborDist < aStar.distance[neighborCoord] {
				aStar.parent[neighborCoord] = current
				aStar.distance[neighborCoord] = neighborDist

				if !aStar.heap.Contains(neighborCoord) {
					aStar.heap.Insert(neighborCoord, neighborDist+T(aStar.heuristic(neighborCoord, end)))
				} else {
					aStar.heap.DecreaseKey(neighborCoord, neighborDist+T(aStar.heuristic(neighborCoord, end)))
				}
			}
		}
	}

	return []domain.Coordinate{}
}

func (aStar *AStarSolver[T]) Initialize(maze *domain.Maze) {
	aStar.maze = maze
	aStar.heap = algorithm.NewHeap[domain.Coordinate, T]()

	aStar.parent = make(map[domain.Coordinate]domain.Coordinate)
	aStar.distance = make(map[domain.Coordinate]T)
	aStar.closedSet = make(map[domain.Coordinate]bool)
}

func (aStar *AStarSolver[T]) isVisited(neighborCoord domain.Coordinate) bool {
	value, ok := aStar.closedSet[neighborCoord]
	return value && ok
}

func (aStar *AStarSolver[T]) inBounds(coordinate domain.Coordinate) bool {
	return coordinate.Column() >= 0 && coordinate.Column() < aStar.maze.Width() &&
		coordinate.Row() >= 0 && coordinate.Row() < aStar.maze.Height()
}

func (aStar *AStarSolver[T]) Neighbors(coordinate domain.Coordinate) []*domain.Cell {
	var neighbors []*domain.Cell

	distance := []domain.Coordinate{
		domain.NewCoordinate(-1, 0),
		domain.NewCoordinate(0, 1),
		domain.NewCoordinate(1, 0),
		domain.NewCoordinate(0, -1),
	}

	current := aStar.maze.Cell(coordinate)

	for i, dist := range distance {
		NeighborCoordinate := domain.Sum(coordinate, dist)
		if aStar.inBounds(NeighborCoordinate) && !current.Walls()[i] {
			neighbors = append(neighbors, aStar.maze.Cell(NeighborCoordinate))
		}
	}

	return neighbors
}
