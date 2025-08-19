package generators

import (
	domain "maze/internal/domain"
	algorithm "maze/pkg/algorithm"
	random "maze/pkg/random"
)

type PrimGenerator struct {
	maze      *domain.Maze
	heap      algorithm.Heap[domain.Coordinate, int]
	weights   []int
	isVisited map[domain.Coordinate]bool
}

func (generator *PrimGenerator) Generate(height, width int) domain.Maze {
	generator.maze = domain.NewMaze(height, width)
	generator.Initialize()

	randomCell := generator.randomCoordinate(generator.maze)
	generator.isVisited[randomCell] = true
	generator.heap.Insert(randomCell, generator.weight(randomCell))

	for !generator.heap.Empty() {
		cell, _ := generator.heap.Min()

		neighbors := generator.unvisitedNeighbor(cell)

		if len(neighbors) == 0 {
			generator.heap.ExtractMin()
		} else {
			randomNeighbor := random.Element(neighbors)
			generator.isVisited[randomNeighbor] = true
			generator.heap.Insert(randomNeighbor, generator.weight(randomNeighbor))
			domain.DeleteWall(generator.maze, cell, randomNeighbor)
		}
	}

	return *generator.maze
}

func (generator *PrimGenerator) randomWeights() []int {
	height := generator.maze.Height()
	width := generator.maze.Width()

	weights := make([]int, width*height)
	for i := 0; i < width*height; i++ {
		weights[i] = i
	}

	weights = algorithm.Shuffle(weights)

	return weights
}

func (generator *PrimGenerator) Initialize() {
	generator.isVisited = make(map[domain.Coordinate]bool)
	generator.heap = algorithm.NewHeap[domain.Coordinate, int]()
	generator.weights = generator.randomWeights()
	generator.isVisited = make(map[domain.Coordinate]bool)

	for i := 0; i < generator.maze.Height(); i++ {
		for j := 0; j < generator.maze.Width(); j++ {
			generator.isVisited[domain.NewCoordinate(i, j)] = false
		}
	}
}

func (generator *PrimGenerator) randomCoordinate(maze *domain.Maze) domain.Coordinate {
	i := random.Index(maze.Grid)
	j := random.Index(maze.Grid[i])

	return domain.NewCoordinate(int(i), int(j))
}

func (generator *PrimGenerator) unvisitedNeighbor(coordinate domain.Coordinate) []domain.Coordinate {
	neighbors := []domain.Coordinate{}
	for _, neighbor := range generator.neighbors(coordinate) {
		if !generator.isVisited[neighbor] {
			neighbors = append(neighbors, neighbor)
		}
	}

	return neighbors
}

func (generator *PrimGenerator) weight(coordinate domain.Coordinate) int {
	return generator.weights[coordinate.Row()*generator.maze.Width()+coordinate.Column()]
}

func inBounds(coordinate domain.Coordinate, width, height int) bool {
	return coordinate.Column() >= 0 && coordinate.Column() < width &&
		coordinate.Row() >= 0 && coordinate.Row() < height
}

func (generator *PrimGenerator) neighbors(coordinate domain.Coordinate) []domain.Coordinate {
	var neighbors []domain.Coordinate

	distance := []domain.Coordinate{
		domain.NewCoordinate(-1, 0),
		domain.NewCoordinate(0, 1),
		domain.NewCoordinate(1, 0),
		domain.NewCoordinate(0, -1),
	}

	for _, dist := range distance {
		NeighborCoordinate := domain.Sum(coordinate, dist)
		if inBounds(NeighborCoordinate, generator.maze.Width(), generator.maze.Height()) {
			neighbors = append(neighbors, NeighborCoordinate)
		}
	}

	return neighbors
}
