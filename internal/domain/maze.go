package domain

import algorithm "maze/pkg/algorithm"

type Maze struct {
	width, height int
	Grid          [][]Cell
}

func NewMaze(height, width int) *Maze {
	maze := &Maze{
		width:  width,
		height: height,
		Grid:   make([][]Cell, height),
	}

	maze.FillMaze()

	return maze
}

func (maze *Maze) FillMaze() {
	LCG := algorithm.NewLCG()

	for i := range maze.Grid {
		maze.Grid[i] = make([]Cell, maze.width)

		for j := range maze.Grid[i] {
			maze.Grid[i][j] = NewCell(i, j, LCG.Next())
		}
	}
}

func (maze *Maze) Cell(coordinate Coordinate) *Cell {
	return &maze.Grid[coordinate.Row()][coordinate.Column()]
}

func (maze *Maze) Width() int {
	return maze.width
}

func (maze *Maze) Height() int {
	return maze.height
}
