package application

import (
	domain "MazeApp/internal/domain"
	"fmt"
)

type MazeApp struct {
	Generator domain.Generator
	Solver    domain.Solver
	Renderer  domain.Renderer
}

func (app MazeApp) Run(width, height int, start, end domain.Coordinate) {
	maze := app.Generator.Generate(height, width)

	path := app.Solver.Solve(maze, start, end)

	if path != nil {
		fmt.Println(app.Renderer.Render(maze))
		fmt.Println(app.Renderer.RenderWithPath(maze, path))
	} else {
		fmt.Println(app.Renderer.Render(maze))
	}
}
