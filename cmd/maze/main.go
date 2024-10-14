package main

import (
	application "MazeApp/internal/application"
	generators "MazeApp/internal/application/generators"
	renderers "MazeApp/internal/application/renderers"
	solvers "MazeApp/internal/application/solvers"
	input "MazeApp/internal/infrastructure/input"
)

func main() {
	width, height := input.GetSize()
	start, end := input.GetRouteDetails(height, width)

	generator := &generators.PrimGenerator{}
	solver := &solvers.AStarSolver[int]{}
	renderer := renderers.SimpleRenderer{}

	app := application.MazeApp{
		Generator: generator,
		Solver:    solver,
		Renderer:  renderer,
	}

	app.Run(width, height, start, end)
}
