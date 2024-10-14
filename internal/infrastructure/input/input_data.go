package input

import (
	domain "MazeApp/internal/domain"
	"fmt"
)

func enterNumber() int {
	var number int

	for {
		number, err := fmt.Scan(&number)
		if err != nil || number < 0 {
			var EndString string

			if _, err := fmt.Scan(&EndString); err != nil {
				fmt.Print()
			}

			fmt.Println("Это не целое положительное число")
			fmt.Print("Введите снова: ")
		} else {
			break
		}
	}

	return number
}

func enterCoordinate(height, width int) (row, column int) {
	for {
		column = enterNumber()
		row = enterNumber()

		if row > height || column > width {
			fmt.Println("Неверные координаты")
			fmt.Print("Введите снова: ")
		} else {
			break
		}
	}

	return row - 1, column - 1
}

func GetSize() (height, width int) {
	fmt.Print("Введите высоту лабиринта: ")

	height = enterNumber()

	fmt.Print("Введите ширину лабиринта: ")

	width = enterNumber()

	return width, height
}

func GetRouteDetails(height, width int) (start, end domain.Coordinate) {
	fmt.Print("Введите координаты начала маршрута (X, Y): ")

	start = domain.NewCoordinate(enterCoordinate(height, width))

	fmt.Print("Введите координаты конца маршрута (X, Y): ")

	end = domain.NewCoordinate(enterCoordinate(height, width))

	return start, end
}
