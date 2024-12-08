package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type Direction struct {
	x int
	y int
}

type Point struct {
	x int
	y int
}

var directions = []Direction{
	{x: 1, y: 0},
	{x: 1, y: 1},
	{x: 0, y: 1},
	{x: -1, y: 1},
	{x: -1, y: 0},
	{x: -1, y: -1},
	{x: 0, y: -1},
	{x: 1, y: -1},
}

func get_input() [][]string {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := [][]string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, strings.Split(scanner.Text(), ""))
	}
	return input
}

func count_pattern_point(input [][]string, point Point, pattern []string) int {
	// First character does not match, nothing left to check
	if input[point.y][point.x] != pattern[0] {
		return 0
	}

	count := 0
	for _, direction := range directions {
		matches := true
		for j, char := range pattern {
			input_x := point.x + (direction.x * j)
			input_y := point.y + (direction.y * j)
			if input_x < 0 || input_y < 0 {
				matches = false
				break
			}
			if input_y > (len(input) - 1) {
				matches = false
				break
			}
			if input_x > (len(input[input_y]) - 1) {
				matches = false
				break
			}
			// log.Printf("input char[%v, %v]: %v, compare char: %v", input_y, input_x, input[input_y][input_x], char)
			if input[input_y][input_x] != char {
				matches = false
				break
			}
		}

		if matches {
			// log.Printf("direction: %v; point: %v", direction, point)
			count += 1
		}
	}
	return count
}

func count_pattern(input [][]string, pattern []string) int {
	count := 0
	for y, chars := range input {
		for x := range chars {
			count += count_pattern_point(input, Point{x: x, y: y}, pattern)
		}
	}
	return count
}

func main() {
	input := get_input()
	count := count_pattern(input, []string{"X", "M", "A", "S"})
	log.Printf("%v", count)
}
