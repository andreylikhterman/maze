package renderer

import (
	"maze/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRender_EmptyMaze(t *testing.T) {
	maze := domain.NewMaze(0, 0)
	renderer := SimpleRenderer{}

	result := renderer.Render(*maze)

	assert.Equal(t, "", result)
}

func TestRenderWithPath_SingleCellMazeWithPath(t *testing.T) {
	maze := domain.NewMaze(1, 1)
	path := []domain.Coordinate{domain.NewCoordinate(0, 0)}
	renderer := SimpleRenderer{}

	expected := "+-------+\n" +
		"|       |\n" +
		"|  📍   |\n" +
		"|       |\n" +
		"+-------+\n"

	result := renderer.RenderWithPath(*maze, path)

	assert.Equal(t, expected, result)
}

func TestRenderWithPath_MultipleCellsMazeWithPath(t *testing.T) {
	maze := domain.NewMaze(1, 2)
	path := []domain.Coordinate{domain.NewCoordinate(0, 0), domain.NewCoordinate(0, 1)}
	renderer := SimpleRenderer{}

	expected := "+-------+-------+\n" +
		"|       |       |\n" +
		"|  📍   |  📍   |\n" +
		"|       |       |\n" +
		"+-------+-------+\n"

	result := renderer.RenderWithPath(*maze, path)

	assert.Equal(t, expected, result)
}
