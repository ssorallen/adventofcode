package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

func Abs[T constraints.Integer](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func sort_asc(i, j int) int {
	return i - j
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	left_values := []int{}
	right_values := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		values := strings.Split(scanner.Text(), "   ")
		left_value, err := strconv.Atoi(values[0])
		if err != nil {
			panic(err)
		}

		right_value, err := strconv.Atoi(values[1])
		if err != nil {
			panic(err)
		}

		left_values = append(left_values, left_value)
		right_values = append(right_values, right_value)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	slices.SortFunc(left_values, sort_asc)
	slices.SortFunc(right_values, sort_asc)

	total_distance := 0
	for i, left_value := range left_values {
		right_value := right_values[i]
		total_distance += Abs(left_value - right_value)
	}

	fmt.Printf("%v", total_distance)
}
