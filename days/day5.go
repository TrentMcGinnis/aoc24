package days

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/trentmcginnis/aoc24/utils"
)

type Page struct {
	num    int64
	before []*Page
	after  []*Page
}

type Manual struct {
	nums []int64
}

func showPageMap(pageMap *map[int64]*Page) {
	for key, value := range *pageMap {
		befores := []int64{}
		afters := []int64{}
		for _, b := range value.before {
			befores = append(befores, b.num)
		}
		for _, a := range value.after {
			afters = append(afters, a.num)
		}
		fmt.Printf("%d - > before: %+v after: %+v\n", key, befores, afters)
	}
}

func sortManual(manual *Manual, pageMap *map[int64]*Page) {
	count := 0
	for i := 0; i < len(manual.nums); i++ {
		if count > 100 {
			fmt.Println("BREAKING")
			break
		}
		redo := false
		if ref, ok := (*pageMap)[manual.nums[i]]; ok {
			for j := 0; j < len(manual.nums); j++ {
				if redo {
					break
				}
				checkingNum := manual.nums[j]
				if j < i {
					// Check that checking num is not in current nums befores list
					for k := 0; k < len(ref.before); k++ {
						if checkingNum == ref.before[k].num {
							// swap
							manual.nums[j] = manual.nums[i]
							manual.nums[i] = checkingNum
							//fmt.Printf("SWAPPING: %d and %d\n", manual.nums[j], manual.nums[i])
							redo = true
							break
						}
					}
				} else if j > i {
					for k := 0; k < len(ref.after); k++ {
						if checkingNum == ref.after[k].num {
							manual.nums[j] = manual.nums[i]
							manual.nums[i] = checkingNum
							//fmt.Printf("SWAPPING: %d and %d\n", manual.nums[j], manual.nums[i])
							redo = true
							break
						}
					}
				}
			}
		}
		if redo {
			i = 0
			count += 1
		}
	}
}

func Day5() {
	pageMap := make(map[int64]*Page)
	lines := utils.GetFile("data/day5/data")
	flag := false
	manuals := []Manual{}
	for _, line := range lines {
		if len(line) == 0 {
			flag = true
			continue
		}
		if flag {
			manual := Manual{}
			nums := strings.Split(line, ",")
			for _, num := range nums {
				parsedNum, _ := strconv.ParseInt(num, 10, 64)
				manual.nums = append(manual.nums, parsedNum)
			}
			manuals = append(manuals, manual)
		} else {
			nums := strings.Split(line, "|")
			//fmt.Printf("PARSING: %v\n", nums)
			var stored *Page
			for i, num := range nums {
				parsedNum, _ := strconv.ParseInt(num, 10, 64)
				var pageRef *Page
				if ref, ok := pageMap[parsedNum]; !ok {
					page := Page{num: parsedNum}
					pageRef = &page
					pageMap[parsedNum] = &page
					if i == 0 {
						stored = &page
					}
				} else {
					pageRef = ref
					if i == 0 {
						stored = ref
					}
				}
				if i == 1 {
					pageRef.after = append(pageRef.after, stored)
					stored.before = append(stored.before, pageRef)
				}
			}
		}
	}
	goodManuals := []*Manual{}
	badManuals := []*Manual{}
	for _, manual := range manuals {
		good := true
		for i := 0; i < len(manual.nums); i++ {
			if !good {
				break
			}
			if pageRef, ok := pageMap[manual.nums[i]]; ok {
				for j := 0; j < len(manual.nums); j++ {
					if !good {
						break
					}
					if j < i {
						// Check that checking num is not in current nums befores list
						checkingNum := manual.nums[j]
						for k := 0; k < len(pageRef.before); k++ {
							if checkingNum == pageRef.before[k].num {
								good = false
								break
							}
						}
					} else if j > i {
						checkingNum := manual.nums[j]
						for k := 0; k < len(pageRef.after); k++ {
							if checkingNum == pageRef.after[k].num {
								good = false
								break
							}
						}
					}
				}
			}
		}
		if good {
			goodManuals = append(goodManuals, &manual)
			//fmt.Printf("G: %+v\n", manual)
		} else {
			badManuals = append(badManuals, &manual)
			//fmt.Printf("B: %+v\n", manual)
		}
	}
	var partOneSum int64 = 0
	for _, man := range goodManuals {
		partOneSum += man.nums[len(man.nums)/2]
	}
	var partTwoSum int64 = 0
	for _, man := range badManuals {
		sortManual(man, &pageMap)
		partTwoSum += man.nums[len(man.nums)/2]
	}
	fmt.Printf("PART ONE: %d\n", partOneSum)
	fmt.Printf("PART TWO: %d\n", partTwoSum)
}
