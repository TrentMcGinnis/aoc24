package days

import (
	"fmt"
	"github.com/trentmcginnis/aoc24/utils"
	"strconv"
	"strings"
)

func Day3() {
	data := utils.GetFile("data/day3/data")
	potentialOps := []string{}
	string := ""
	for _, line := range data {
		string += line
		for i := 0; i < len(line); i++ {
			char := line[i]
			if char == 'm' {
				end := 12
				if len(line)-i < 12 {
					end = len(line) - i
				}
				for j := 7; j < end; j++ {
					checkChar := line[i+j]
					if checkChar == ')' {
						potentialOps = append(potentialOps, line[i:i+j+1])
						break
					}
				}
			} else if char == 'd' {
				end := 7
				if len(line)-i < 7 {
					end = len(line) - i
				}
				for j := 3; j < end; j++ {
					checkChar := line[i+j]
					if checkChar == ')' {
						potentialOps = append(potentialOps, line[i:i+j+1])
						break
					}
				}
			}
		}
	}
	partOneSum := 0
	partTwoSum := 0
	enabled := true
	for _, p := range potentialOps {
		fmt.Printf("%+v\n", p)
		if p[0] == 'd' {
			if len(p) >= 4 {
				if p[:4] == "do()" {
					enabled = true
				} else if len(p) >= 7 && p[:7] == "don't()" {
					enabled = false
				}
			}
		} else {
			nums := strings.Split(p, ",")
			fmt.Printf("%+v\n", nums)
			if len(nums) == 2 && len(nums[0]) >= 4 && nums[0][:4] == "mul(" && len(nums[0])+len(nums[1]) >= 7 {
				numOne, errOne := strconv.ParseInt(nums[0][4:], 10, 64)
				numTwo, errTwo := strconv.ParseInt(nums[1][:len(nums[1])-1], 10, 64)
				fmt.Printf("%+v %+v\n\n", numOne, numTwo)
				if errOne == nil && errTwo == nil {
					product := numOne * numTwo
					partOneSum += int(product)
					if enabled {
						partTwoSum += int(product)
					}
				}
			}
		}
	}
	fmt.Printf("Part One: %d\n", partOneSum)
	fmt.Printf("Part Two: %d\n", partTwoSum)
}
