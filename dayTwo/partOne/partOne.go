package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../input.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	totalValid := 0

	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Fields(line)
		increasing := false
		valid := 0
		for n := range len(nums) {
			if n+1 != len(nums) {

				curr, err := strconv.Atoi(nums[n])
				if err != nil {
					panic(err)
				}
				plusOne := n + 1
				next, err := strconv.Atoi(nums[plusOne])
				if err != nil {
					fmt.Println(nums[n])
					fmt.Println(nums)
					panic(err)
				}

				if n == 0 {
					increasing = curr < next
				}

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
					totalValid += 1
				}
			}
		}
	}

	fmt.Println(totalValid)
}
