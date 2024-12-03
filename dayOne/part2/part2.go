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
	file, err := os.Open("./input.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	leftList := make([]int, 0)
	rightList := make(map[string]int)

	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Fields(line)
		leftList = append(leftList, parseNumber(nums[0]))
		rightList[nums[1]] += 1
		fmt.Println(nums[1])
	}
	fmt.Println(rightList)
	slices.Sort(leftList)

	diff := 0

	for i := range len(leftList) {
		valA := leftList[i]
		valB := rightList[strconv.Itoa(valA)]
		calc := valA * valB
		// fmt.Println(valB)
		diff += calc
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
