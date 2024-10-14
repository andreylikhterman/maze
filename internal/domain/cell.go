package domain

const (
	EMPTY = 3
	COIN  = 1
	SAND  = 4
	SWAMP = 6
)

type Surface = int

type Wall = bool

type Cell struct {
	coordinate Coordinate
	walls      [4]Wall
	surface    Surface
}

func NewCell(row, column int, randomNumber int64) Cell {
	surfaces := []Surface{EMPTY, EMPTY, EMPTY, COIN, SAND, SWAMP}
	randomSurface := surfaces[randomNumber%int64(len(surfaces))]
	cell := Cell{
		coordinate: Coordinate{
			row:    row,
			column: column,
		},

		walls:   [4]Wall{true, true, true, true},
		surface: randomSurface,
	}

	return cell
}

func (cell *Cell) Coordinate() Coordinate {
	return cell.coordinate
}

func (cell *Cell) Walls() [4]Wall {
	return cell.walls
}

func (cell *Cell) Surface() Surface {
	return cell.surface
}

func DeleteWall(maze *Maze, start, end Coordinate) {
	if start.Row() == end.Row() {
		deleteVerticalWall(maze, start, end)
	} else {
		deleteHorizontalWall(maze, start, end)
	}
}

func deleteHorizontalWall(maze *Maze, start, end Coordinate) {
	cell1 := maze.Cell(start)
	cell2 := maze.Cell(end)

	if start.Row()+1 == end.Row() {
		cell1.walls[2] = false
		cell2.walls[0] = false
	} else {
		cell1.walls[0] = false
		cell2.walls[2] = false
	}
}

func deleteVerticalWall(maze *Maze, start, end Coordinate) {
	cell1 := maze.Cell(start)
	cell2 := maze.Cell(end)

	if start.Column()+1 == end.Column() {
		cell1.walls[1] = false
		cell2.walls[3] = false
	} else {
		cell1.walls[3] = false
		cell2.walls[1] = false
	}
}
