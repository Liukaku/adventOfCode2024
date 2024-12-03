package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	totalValid := 0

	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Fields(line)

		// slices.Sort(copyArr)
		length := len(nums) - 1
		increasing := parseInt(nums[0]) < parseInt(nums[length])

		if isValid(nums, increasing) {
			totalValid += 1
		} else {
			for k := range len(nums) {
				copyArr := make([]string, len(nums))
				copy(copyArr, nums)
				newSlice := make([]string, len(copyArr)-1)
				newSlice = append(copyArr[:k], copyArr[k+1:]...)
				length := len(newSlice) - 1
				increasing := parseInt(newSlice[0]) < parseInt(newSlice[length])
				if isValid(newSlice, increasing) {
					totalValid += 1

					break
				} else {

				}
			}
		}

	}

	fmt.Println(totalValid)
}

func isValid(nums []string, increasing bool) bool {
	valid := 0
	success := false
	for n := range len(nums) {
		if n+1 != len(nums) {

			curr := parseInt(nums[n])
			plusOne := n + 1
			next := parseInt(nums[plusOne])

			if increasing {
				if next-curr <= 3 && next-curr > 0 {
					valid += 1
				}
			} else {
				if curr-next <= 3 && curr-next > 0 {
					valid += 1
				}
			}
		} else {
			if valid+1 == len(nums) {
				success = true
			}
		}
	}
	return success
}

func parseInt(val string) int {
	next, err := strconv.Atoi(val)
	if err != nil {
		panic(err)
	}
	return next
}
