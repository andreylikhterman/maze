package domain

type Generator interface {
	Generate(height, width int) Maze
}

type Solver interface {
	Solve(maze Maze, start, end Coordinate) []Coordinate
}

type Renderer interface {
	Render(maze Maze) string
	RenderWithPath(maze Maze, path []Coordinate) string
}
