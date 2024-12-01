package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func get_values() [][]int {
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

	return_val := [][]int{left_values, right_values}
	return return_val
}

func get_counts(values []int) map[int]int {
	counts_map := make(map[int]int)
	for _, value := range values {
		counts_map[value] += 1
	}
	return counts_map
}

func main() {
	values := get_values()
	right_values := values[1]
	right_value_counts := get_counts(right_values)

	similarity_score := 0
	left_values := values[0]
	for _, value := range left_values {
		if count, ok := right_value_counts[value]; ok {
			similarity_score += value * count
		}
	}

	fmt.Printf("%v", similarity_score)
}
