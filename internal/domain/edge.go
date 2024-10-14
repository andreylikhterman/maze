package domain

type Edge struct {
	start, end *Cell
}

func (edge Edge) Start() *Cell {
	return edge.start
}

func (edge Edge) End() *Cell {
	return edge.end
}

func NewEdge(maze *Maze, start, end Coordinate) Edge {
	edge := Edge{
		start: maze.Cell(NewCoordinate(start.Row(), start.Column())),
		end:   maze.Cell(NewCoordinate(end.Row(), end.Column())),
	}

	return edge
}
