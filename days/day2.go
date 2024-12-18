package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isReportSafe(nums []int64) bool {
	increasing, decreasing := false, false
	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		if diff == 0 {
			return false
		} else if diff < 0 {
			decreasing = true
		} else {
			increasing = true
		}

		if increasing && decreasing {
			return false
		}

		if diff > 3 || diff < -3 {
			return false
		}
	}
	return true
}

func isReportSafeWithDelete(line []int64, delete int) bool {
	copied := make([]int64, len(line))
	copy(copied, line)

	if delete == len(line)-1 {
		copied = copied[:delete]
	} else {
		copied = append(copied[:delete], copied[delete+1:]...)
	}
	return isReportSafe(copied)
}

func checkReportWithDelete(line []int64) bool {
	for i := 0; i < len(line); i++ {
		safe := isReportSafeWithDelete(line, i)
		if safe {
			return true
		}
	}
	return false
}

func Day2() {
	file, err := os.Open("data/day2/data")
	if err != nil {
		fmt.Printf("Error opening file: %s\n", err)
	}

	lines := [][]int64{}
	scanner := bufio.NewScanner(file)
	for more := scanner.Scan(); more; more = scanner.Scan() {
		line := []int64{}
		characters := strings.Split(scanner.Text(), " ")
		for i := 0; i < len(characters); i++ {
			if num, err := strconv.ParseInt(characters[i], 10, 64); err == nil {
				line = append(line, num)
			}
		}
		lines = append(lines, line)
	}
	partOneSafe := 0
	partTwoSafe := 0
	for _, line := range lines {
		safety := isReportSafe(line)
		if safety {
			partOneSafe += 1
		} else if checkReportWithDelete(line) {
			partTwoSafe += 1
		}
	}
	fmt.Printf("PART 1: %d\n", partOneSafe)
	fmt.Printf("PART 2: %d\n", partOneSafe+partTwoSafe)
}
