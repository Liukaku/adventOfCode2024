package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	leftList := make([]int, 0)
	rightList := make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Fields(line)
		leftList = append(leftList, parseNumber(nums[0]))
		rightList = append(rightList, parseNumber(nums[1]))
	}
	slices.Sort(leftList)
	slices.Sort(rightList)

	diff := 0

	for i := range len(leftList) {
		if leftList[i] > rightList[i] {
			diff += leftList[i] - rightList[i]
		} else {
			diff += rightList[i] - leftList[i]
		}
	}
	fmt.Println(diff)
}

func parseNumber(toParse string) int {
	parsedInt, err := strconv.ParseInt(toParse, 0, 0)
	if err != nil {
		panic(err)
	}
	return int(parsedInt)
}
