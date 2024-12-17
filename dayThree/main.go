package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	// do\(\)\S{1,}mul\((\d{1,3})\,(\d{1,3})\) ??
	// mul\((\d{1,3})\,(\d{1,3})\)|do\(\)|don\'t\(\)
	count := 0
	do := true
	for scanner.Scan() {
		line := scanner.Text()

		r, _ := regexp.Compile(`mul\((\d{1,3})\,(\d{1,3})\)|do\(\)|don\'t\(\)`)
		matches := r.FindAllStringSubmatch(line, -1)
		for _, v := range matches {
			fmt.Println(v[0] == "do()")
			if v[0] == "don't()" {
				do = false
			} else if v[0] == "do()" {
				do = true
			} else if do == true {
				groups := r.FindStringSubmatch(v[0])
				a, _ := strconv.Atoi(groups[1])
				b, _ := strconv.Atoi(groups[2])
				count += (a * b)
			}
		}

	}
	fmt.Printf("Total Count: %s \n", strconv.Itoa(count))
}

func calcTotal(str string) int {
	r, _ := regexp.Compile("(mul\\((\\d{1,})\\,(\\d{1,99})\\))")
	matches := r.FindAllStringSubmatch(str, -1)
	count := 0
	for _, v := range matches {
		// fmt.Println(v[0])
		groups := r.FindStringSubmatch(v[0])
		a, _ := strconv.Atoi(groups[2])
		b, _ := strconv.Atoi(groups[3])
		count += (a * b)
	}
	return count
}
