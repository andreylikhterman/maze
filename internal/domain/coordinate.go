package domain

type Coordinate struct {
	row    int
	column int
}

func (coordinate Coordinate) Row() int {
	return coordinate.row
}

func (coordinate Coordinate) Column() int {
	return coordinate.column
}

func NewCoordinate(row, column int) Coordinate {
	return Coordinate{
		row:    row,
		column: column,
	}
}

func Sum(coordinates ...Coordinate) (end Coordinate) {
	for _, coordinate := range coordinates {
		end.column += coordinate.Column()
		end.row += coordinate.Row()
	}

	return end
}
