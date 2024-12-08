package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type Point struct {
	x int
	y int
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
	// First character does not match center; nothing left to check
	if input[point.y][point.x] != pattern[1] {
		return 0
	}

	if point.y < 1 || point.y > len(input)-2 {
		return 0
	}

	if point.x < 1 || point.x > len(input[point.y])-2 {
		return 0
	}

	if ((input[point.y-1][point.x-1] == pattern[0] && input[point.y+1][point.x+1] == pattern[2]) ||
		(input[point.y-1][point.x-1] == pattern[2] && input[point.y+1][point.x+1] == pattern[0])) &&
		((input[point.y+1][point.x-1] == pattern[0] && input[point.y-1][point.x+1] == pattern[2]) ||
			(input[point.y+1][point.x-1] == pattern[2] && input[point.y-1][point.x+1] == pattern[0])) {
		return 1
	}

	return 0
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
	count := count_pattern(input, []string{"M", "A", "S"})
	log.Printf("%v", count)
}
