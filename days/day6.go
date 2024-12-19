package days

import (
	"fmt"

	"github.com/trentmcginnis/aoc24/utils"
)

type Cell struct {
	value   byte
	kind    int
	visited bool
	x       int
	y       int
	ogX     int
	ogY     int
	guard   bool
}

func displayCells(cells [][]*Cell, direction int) {
	for y := 0; y < len(cells); y++ {
		for x := 0; x < len(cells[y]); x++ {
			cell := cells[y][x]
			printVal := cell.value
			if cell.visited {
				printVal = 'X'
			}
			if cell.guard {
				switch direction {
				case 0:
					printVal = '^'
					break
				case 1:
					printVal = '>'
					break
				case 2:
					printVal = 'v'
					break
				case 3:
					printVal = '<'
					break
				}
			}
			fmt.Printf("%c ", printVal)
		}
		fmt.Println()
	}
	fmt.Println("___________________")
}

func traverseLine(cells [][]*Cell, guard *Cell, direction int) {
	//displayCells(cells, direction)
	var nextCell *Cell
	switch direction {
	case 0:
		if nextY := guard.y - 1; nextY > -1 {
			nextCell = cells[nextY][guard.x]
		}
		break
	case 1:
		if nextX := guard.x + 1; nextX < len(cells[0]) {
			nextCell = cells[guard.y][nextX]
		}
		break
	case 2:
		if nextY := guard.y + 1; nextY < len(cells) {
			nextCell = cells[nextY][guard.x]
		}
		break
	case 3:
		if nextX := guard.x - 1; nextX > -1 {
			nextCell = cells[guard.y][nextX]
		}
		break
	}
	if nextCell != nil {
		cellRef := cells[guard.y][guard.x]
		//fmt.Printf("CURRENT CELL: %d,%d -> %+v\n", cellRef.x, cellRef.y, cellRef)
		//fmt.Printf("NEXT CELL: %d,%d -> %+v\n", nextCell.x, nextCell.y, nextCell)
		if nextCell.kind == 1 {
			direction = (direction + 1) % 4
		} else {
			cellRef.visited = true
			cellRef.kind = 0
			cellRef.guard = false
			guard.x = nextCell.ogX
			guard.y = nextCell.ogY
			nextCell.guard = true
		}
		traverseLine(cells, guard, direction)
	}
}

func Day6() {
	cells := [][]*Cell{}
	lines := utils.GetFile("data/day6/data")
	startX := 0
	startY := 0
	for y := 0; y < len(lines); y++ {
		cellLine := []*Cell{}
		for x := 0; x < len(lines[y]); x++ {
			var kind int
			char := lines[y][x]
			switch char {
			case '.':
				kind = 0
				break
			case '#':
				kind = 1
				break
			case '^':
				kind = 2
				startX = x
				startY = y
				break
			}
			cell := Cell{
				kind:    kind,
				visited: false,
				value:   char,
				x:       x,
				y:       y,
				guard:   kind == 2,
				ogX:     x,
				ogY:     y,
			}
			cellLine = append(cellLine, &cell)
		}
		cells = append(cells, cellLine)
	}
	guard := cells[startY][startX]
	direction := 0
	traverseLine(cells, guard, direction)
	count := 0
	for y := 0; y < len(cells); y++ {
		for x := 0; x < len(cells[y]); x++ {
			cell := cells[y][x]
			if cell.visited || cell.guard {
				count += 1
			}
		}
	}
	fmt.Printf("PART ONE: %d\n", count)
}
