package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func get_input() string {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func find_instructions(input string) [][]string {
	re := regexp.MustCompile(`(don't)\(\)|(do)\(\)|(mul)\((\d{1,3}),(\d{1,3})\)`)
	return re.FindAllStringSubmatch(input, -1)
}

func run_instructions(instructions [][]string) int {
	enabled := true
	sum := 0
	for _, instruction := range instructions {
		inst := instruction[1]
		if len(inst) == 0 {
			inst = instruction[2]
		}
		if len(inst) == 0 {
			inst = instruction[3]
		}

		switch inst {
		case "do":
			enabled = true
		case "don't":
			enabled = false
		case "mul":
			if !enabled {
				continue
			}
			op1, err := strconv.Atoi(instruction[4])
			if err != nil {
				panic(err)
			}
			op2, err := strconv.Atoi(instruction[5])
			if err != nil {
				panic(err)
			}
			sum += op1 * op2
		}
	}
	return sum
}

func main() {
	input := get_input()
	instructions := find_instructions(input)
	// fmt.Printf("%v", instructions)
	fmt.Printf("%v", run_instructions(instructions))
}
