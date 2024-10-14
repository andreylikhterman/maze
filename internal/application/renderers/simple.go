package renderer

import (
	domain "MazeApp/internal/domain"
)

const (
	emptyString = "       "
)

type SimpleRenderer struct{}

func (simple SimpleRenderer) Render(maze domain.Maze) string {
	var result string

	if maze.Height() != 0 {
		result += renderRowTop(maze.Grid[0])
	}

	for i := 0; i < maze.Height(); i++ {
		if len(maze.Grid[i]) > 0 {
			result += renderRowMiddle(maze.Grid[i])
			result += renderRowBottom(maze.Grid[i])
		}
	}

	return result
}

func renderRowTop(row []domain.Cell) string {
	var result string

	for _, cell := range row {
		if cell.Walls()[0] {
			result += "+-------"
		} else {
			result += "+       "
		}
	}

	result += "+\n"

	return result
}

func renderRowMiddle(row []domain.Cell) string {
	var result string

	for h := 0; h < 3; h++ {
		for _, cell := range row {
			result += renderCellMiddle(cell, h)
		}

		result += "|\n"
	}

	return result
}

func renderCellMiddle(cell domain.Cell, row int) string {
	var result string

	if cell.Walls()[3] {
		result += "|"
	} else {
		result += " "
	}

	if row == 1 {
		result += renderSurface(cell)
	} else {
		result += emptyString
	}

	return result
}

func renderRowBottom(row []domain.Cell) string {
	var result string

	for _, cell := range row {
		if cell.Walls()[2] {
			result += "+-------"
		} else {
			result += "+       "
		}
	}

	result += "+\n"

	return result
}

func renderSurface(cell domain.Cell) string {
	switch cell.Surface() {
	case domain.COIN:
		return "   💵  "
	case domain.SAND:
		return "   🐫  "
	case domain.SWAMP:
		return "   🐸  "
	default:
		return emptyString
	}
}

func (simple SimpleRenderer) RenderWithPath(maze domain.Maze, path []domain.Coordinate) string {
	var result string

	result += renderRowTop(maze.Grid[0])

	for i := 0; i < maze.Height(); i++ {
		result += renderRowMiddleWithPath(maze.Grid[i], path)
		result += renderRowBottom(maze.Grid[i])
	}

	return result
}

func renderRowMiddleWithPath(row []domain.Cell, path []domain.Coordinate) string {
	var result string

	for h := 0; h < 3; h++ {
		for _, cell := range row {
			result += renderCellMiddleWithPath(cell, h, path)
		}

		result += "|\n"
	}

	return result
}

func renderCellMiddleWithPath(cell domain.Cell, row int, path []domain.Coordinate) string {
	var result string

	if cell.Walls()[3] {
		result += "|"
	} else {
		result += " "
	}

	if row == 1 {
		result += renderSurfaceWithPath(cell, path)
	} else {
		result += emptyString
	}

	return result
}

func renderSurfaceWithPath(cell domain.Cell, path []domain.Coordinate) string {
	for _, coord := range path {
		if cell.Coordinate() == coord {
			return "  🟥   "
		}
	}

	return renderSurface(cell)
}
