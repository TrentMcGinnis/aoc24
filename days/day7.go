package days

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/trentmcginnis/aoc24/utils"
)

type Expression struct {
	left     interface{}
	right    interface{}
	operator string
	total    int64
}

type Line struct {
	total          int64
	expressionHead *Expression
}

func DisplayExpression(expression *Expression, total string) string {
	switch expression.left.(type) {
	case int64:
		return fmt.Sprintf("(%d %s %d)", expression.left, expression.operator, expression.right)
	case *Expression:
		return fmt.Sprintf("(%s %s %d)", DisplayExpression(expression.left.(*Expression), total), expression.operator, expression.right)
	}

	return ""
}

func ParseExpression(expression *Expression, all *[]int64) []int64 {
	//fmt.Printf("PARSING! %s\n", DisplayExpression(expression, ""))
	returnArgs := []int64{}
	switch expression.left.(type) {
	case int64:
		for i := 0; i < 3; i++ {
			switch i {
			case 0:
				//fmt.Printf("%d + %d\n", expression.left, expression.right)
				returnArgs = append(returnArgs, expression.left.(int64)+expression.right.(int64))
				break

			case 1:
				//fmt.Printf("%d * %d\n", expression.left, expression.right)
				returnArgs = append(returnArgs, expression.left.(int64)*expression.right.(int64))
				break
			case 2:
				//fmt.Printf("%d * %d\n", expression.left, expression.right)
				newInt, _ := strconv.ParseInt(fmt.Sprintf("%d%d", expression.left.(int64), expression.right.(int64)), 10, 64)
				returnArgs = append(returnArgs, int64(newInt))
				break
			}
		}
		return returnArgs
	case *Expression:
		parsed := ParseExpression(expression.left.(*Expression), all)
		newParsed := []int64{}
		for _, par := range parsed {
			newExpression := &Expression{left: par, right: expression.right}
			moreParsed := ParseExpression(newExpression, all)
			newParsed = append(newParsed, moreParsed...)
		}
		*all = append(*all, newParsed...)
		return newParsed
	}

	return returnArgs
}

func Day7() {
	lines := utils.GetFile("data/day7/data")
	parsedLines := []Line{}
	partOneTotal := 0
	for j := 0; j < len(lines); j++ {
		line := lines[j]
		splitLine := strings.Split(line, ":")
		total, _ := strconv.ParseInt(splitLine[0], 10, 64)
		parsedLine := Line{total: total}

		var currentExpression *Expression
		var nextExpression *Expression
		operands := strings.Split(strings.Trim(splitLine[1], " "), " ")
		//fmt.Printf("%d -> %+v\n", total, operands)
		for i := len(operands) - 1; i > 0; i-- {
			operand, _ := strconv.ParseInt(operands[i], 10, 64)
			if nextExpression != nil {
				currentExpression = nextExpression
				currentExpression.right = operand
				nextExpression = nil
			} else {
				currentExpression = &Expression{right: operand, operator: "+"}
				if parsedLine.expressionHead == nil {
					parsedLine.expressionHead = currentExpression
				}
			}
			if i-1 > 0 {
				nextExpression = &Expression{operator: "+"}
				currentExpression.left = nextExpression
			} else {
				nextOperand, _ := strconv.ParseInt(operands[i-1], 10, 64)
				currentExpression.left = nextOperand
			}
		}
		parsedLines = append(parsedLines, parsedLine)
		all := []int64{}
		//fmt.Println(DisplayExpression(parsedLine.expressionHead, ""))
		all = append(all, ParseExpression(parsedLine.expressionHead, &all)...)
		for _, sum := range all {
			//fmt.Printf("%d == %d = %t\n", sum, parsedLine.total, sum == parsedLine.total)
			if sum == parsedLine.total {
				partOneTotal += int(parsedLine.total)
				break
			}
		}
		//fmt.Printf("RUNNING TOTAL: %d\n", partOneTotal)
	}
	fmt.Printf("PART ONE TOTAL: %d\n", partOneTotal)
}
