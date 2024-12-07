package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Report = []Level
type Level = int

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func is_report_safe(report Report, index_to_ignore int) bool {
	prev_level := -1
	report_type := "unknown"
	for i, level := range report {
		if i == index_to_ignore {
			continue
		} else if prev_level == -1 {
			prev_level = level
		} else if level-prev_level == 0 || abs(level-prev_level) > 3 {
			return false
		} else if report_type == "unknown" {
			if level > prev_level {
				report_type = "increasing"
			} else {
				report_type = "decreasing"
			}
		} else if report_type == "increasing" && level < prev_level {
			return false
		} else if report_type == "decreasing" && level > prev_level {
			return false
		}
		prev_level = level
	}
	return true
}

func is_report_mostly_safe(report Report) bool {
	index_to_ignore := -1
	for index_to_ignore < len(report) {
		if is_report_safe(report, index_to_ignore) {
			return true
		}
		index_to_ignore += 1
	}
	return false
}

func count_safe_reports(reports []Report) int {
	count := 0
	for _, report := range reports {
		is_safe := is_report_mostly_safe((report))
		if is_safe {
			count += 1
		}
	}
	return count
}

func get_reports() []Report {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reports := []Report{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		levels_str := strings.Split(scanner.Text(), " ")
		levels_int := make([]int, len(levels_str))
		for i, level_str := range levels_str {
			level_int, err := strconv.Atoi(level_str)
			if err != nil {
				panic(err)
			}
			levels_int[i] = level_int
		}

		reports = append(reports, levels_int)
	}

	return reports
}

func main() {
	reports := get_reports()
	safe_reports_count := count_safe_reports(reports)

	fmt.Printf("%v", safe_reports_count)
}
