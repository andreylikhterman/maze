package main

import (
	application "maze/internal/application"
	generators "maze/internal/application/generators"
	renderers "maze/internal/application/renderers"
	solvers "maze/internal/application/solvers"
	input "maze/internal/infrastructure/input"
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
