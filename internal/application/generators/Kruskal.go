package generators

import (
	domain "maze/internal/domain"
	algorithm "maze/pkg/algorithm"
)

type KruskalGenerator struct{}

func (generator *KruskalGenerator) Generate(height, width int) domain.Maze {
	maze := domain.NewMaze(height, width)
	edges := generator.edges(maze, height, width)
	edges = algorithm.Shuffle(edges)

	disjointSet := algorithm.NewDisjointSet[domain.Coordinate]()

	for row := range height {
		for column := range width {
			disjointSet.MakeSet(domain.NewCoordinate(row, column))
		}
	}

	for _, edge := range edges {
		startCoord := edge.Start().Coordinate()
		endCoord := edge.End().Coordinate()

		if disjointSet.Find(startCoord) != disjointSet.Find(endCoord) {
			domain.DeleteWall(maze, startCoord, endCoord)
			disjointSet.Union(startCoord, endCoord)
		}
	}

	return *maze
}

func (generator *KruskalGenerator) edges(maze *domain.Maze, height, width int) []domain.Edge {
	edges := []domain.Edge{}

	for row := range height {
		for column := range width {
			start := domain.NewCoordinate(row, column)

			if column < width-1 {
				end := domain.NewCoordinate(row, column+1)
				edge := domain.NewEdge(maze, start, end)
				edges = append(edges, edge)
			}

			if row < height-1 {
				end := domain.NewCoordinate(row+1, column)
				edge := domain.NewEdge(maze, start, end)
				edges = append(edges, edge)
			}
		}
	}

	return edges
}
