package days

import (
	"fmt"
	"strings"

	"github.com/trentmcginnis/aoc24/utils"
)

type Node struct {
	x         int
	y         int
	value     string
	left      *Node
	right     *Node
	up        *Node
	down      *Node
	upLeft    *Node
	upRight   *Node
	downLeft  *Node
	downRight *Node
}

func recurseDirection(node *Node, count int, total string, direction uint) string {
	var nextNode *Node
	var directionString string
	if count == 3 {
		return total
	}
	switch direction {
	case 0:
		nextNode = node.upLeft
		directionString = "UP LEFT"
		break
	case 1:
		nextNode = node.up
		directionString = "UP"
		break
	case 2:
		nextNode = node.upRight
		directionString = "UP RIGHT"
		break
	case 3:
		nextNode = node.left
		directionString = "LEFT"
		break
	case 4:
		nextNode = node.right
		directionString = "RIGHT"
		break
	case 5:
		nextNode = node.downLeft
		directionString = "DOWN LEFT"
		break
	case 6:
		nextNode = node.down
		directionString = "DOWN"
		break
	case 7:
		nextNode = node.downRight
		directionString = "DOWN RIGHT"
		break
	}
	if count == 0 && false {
		total = fmt.Sprintf("%s ", directionString) + total
	}
	if nextNode != nil {
		total += nextNode.value

	} else {
		return total
	}
	count += 1
	return recurseDirection(nextNode, count, total, direction)
}

func isXmas(node *Node) bool {
	upLeft := node.upLeft
	upRight := node.upRight
	downLeft := node.downLeft
	downRight := node.downRight
	if upLeft != nil && downRight != nil && downLeft != nil && upRight != nil {
		strOne := upLeft.value + node.value + downRight.value
		strTwo := downLeft.value + node.value + upRight.value
		return (strOne == "MAS" || strOne == "SAM") && (strTwo == "MAS" || strTwo == "SAM")
	} else {
		return false
	}
}

func Day4() {
	lines := utils.GetFile("data/day4/data")
	grid := [][]*Node{}
	for y := 0; y < len(lines); y++ {
		line := lines[y]
		gridLine := []*Node{}
		chars := strings.Split(line, "")
		for x := 0; x < len(chars); x++ {
			node := Node{x: x, y: y, value: chars[x]}
			if x > 0 {
				left := gridLine[x-1]
				node.left = left
				left.right = &node
				if y > 0 {
					upLeft := grid[y-1][x-1]
					node.upLeft = upLeft
					upLeft.downRight = &node
				}
			}
			if y > 0 {
				up := grid[y-1][x]
				node.up = up
				up.down = &node
				if x < len(chars)-1 {
					upRight := grid[y-1][x+1]
					node.upRight = upRight
					upRight.downLeft = &node
				}
			}
			gridLine = append(gridLine, &node)
		}
		grid = append(grid, gridLine)
	}
	allStrings := []string{}
	partTwoCount := 0
	for _, row := range grid {
		for _, col := range row {
			//value := fmt.Sprintf("%d:%d -> %s", col.x, col.y, col.value)
			value := fmt.Sprintf("%s", col.value)
			colStringUpLeft := recurseDirection(col, 0, value, 0)
			colStringUp := recurseDirection(col, 0, value, 1)
			colStringUpRight := recurseDirection(col, 0, value, 2)
			colStringLeft := recurseDirection(col, 0, value, 3)
			colStringRight := recurseDirection(col, 0, value, 4)
			colStringDownLeft := recurseDirection(col, 0, value, 5)
			colStringDown := recurseDirection(col, 0, value, 6)
			colStringDownRight := recurseDirection(col, 0, value, 7)
			//fmt.Printf("%d:%d -> %s|", col.x, col.y, col.value)
			allStrings = append(allStrings, colStringUpLeft, colStringUp, colStringUpRight)
			allStrings = append(allStrings, colStringLeft, colStringRight)
			allStrings = append(allStrings, colStringDownLeft, colStringDown, colStringDownRight)
			if col.value == "A" && isXmas(col) {
				partTwoCount += 1
			}
		}
		//fmt.Println()
	}
	count := 0
	for _, s := range allStrings {
		//fmt.Println(s)
		if len(s) == 4 {
			reverse := ""
			for i := len(s) - 1; i >= 0; i-- {
				reverse += string(s[i])
			}
			if reverse == "XMAS" || s == "XMAS" {
				count += 1
			}
		}
	}
	fmt.Printf("PART ONE: %d\n", count/2)
	fmt.Printf("PART TWO: %d\n", partTwoCount)
}
