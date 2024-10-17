package input

import (
	domain "MazeApp/internal/domain"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type numbers interface {
	int | int16 | int32 | int64 | float32 | float64
}

func checkRange[T numbers](value, start, end T) bool {
	return (value >= start && value <= end)
}

func isValidNumber(input string) bool {
	reg := regexp.MustCompile(`^-?\d+$`)
	return reg.MatchString(input)
}

func isValidCoordinates(input string) bool {
	re := regexp.MustCompile(`^-?\d+\s+-?\d+$`)
	return re.MatchString(input)
}

func enterNumber() int {
	reader := bufio.NewReader(os.Stdin)

	var number int

	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if !isValidNumber(input) {
			fmt.Print("Это не целое число.\nВведите снова: ")
			continue
		}

		number, _ = strconv.Atoi(input)
		if number <= 0 {
			fmt.Print("Это не положительное число.\nВведите снова: ")
			continue
		}

		break
	}

	return number
}

func enterCoordinate(width, height int) (row, column int) {
	reader := bufio.NewReader(os.Stdin)

	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if !isValidCoordinates(input) {
			fmt.Print("Неверный формат координат.\nВведите снова: ")
			continue
		}

		coords := strings.Fields(input)
		column, _ = strconv.Atoi(coords[0])
		row, _ = strconv.Atoi(coords[1])

		if !(checkRange(column, 0, width-1) && checkRange(row, 0, height-1)) {
			fmt.Printf("Координаты вне допустимых границ (X: 0-%d, Y: 0-%d).\nВведите снова: ", width-1, height-1)
			continue
		}

		break
	}

	return row, column
}

func GetSize() (height, width int) {
	fmt.Print("Введите ширину лабиринта: ")

	width = enterNumber()

	fmt.Print("Введите высоту лабиринта: ")

	height = enterNumber()

	return width, height
}

func GetRouteDetails(height, width int) (start, end domain.Coordinate) {
	fmt.Print("Введите координаты начала маршрута (X, Y): ")

	start = domain.NewCoordinate(enterCoordinate(width, height))

	fmt.Print("Введите координаты конца маршрута (X, Y): ")

	end = domain.NewCoordinate(enterCoordinate(width, height))

	return start, end
}
