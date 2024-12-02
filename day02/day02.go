package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func abs_diff(a int, b int) int {
	res := 0
	if b > a {
		res = b - a
	} else {
		res = a - b
	}

	return res
}

func eval_increasing(levels []int, inc bool) bool {
	for i := 1; i < len(levels); i++ {
		if levels[i] > levels[i-1] != inc {
			return false
		}
	}

	return true
}

func eval_minimum(levels []int, minimum int) bool {
	for i := 1; i < len(levels); i++ {
		if abs_diff(levels[i], levels[i-1]) < minimum {
			return false
		}
	}

	return true
}

func eval_maximum(levels []int, maximum int) bool {
	for i := 1; i < len(levels); i++ {
		if abs_diff(levels[i], levels[i-1]) > maximum {
			return false
		}
	}

	return true
}

func is_safe(levels []int) bool {
	inc := levels[1] > levels[0]
	return eval_increasing(levels, inc) && eval_minimum(levels, 1) && eval_maximum(levels, 3)
}

func str_to_int(in []string) []int {
	out := []int{}

	for _, i := range in {
		n, err := strconv.Atoi(i)
		if err != nil {
			log.Fatal("could not parse integer")
		}

		out = append(out, n)
	}

	return out
}

func part_one(levels [][]int) int {
	safe_reports := 0

	for _, levels := range levels {
		if is_safe(levels) {
			safe_reports += 1
		}
	}
	return safe_reports
}

func part_two(levels [][]int) int {
	safe_reports := 0

	for _, levels := range levels {
		safe := is_safe(levels)

		skip := 0
		for !safe && skip < len(levels) {
			skipped := make([]int, len(levels)-1)
			copy(skipped, levels[:skip])
			copy(skipped[skip:], levels[skip+1:])

			safe = is_safe(skipped)

			skip += 1
		}

		if safe {
			safe_reports += 1
		}

	}
	return safe_reports
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	all_levels := [][]int{}

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		levels := str_to_int(parts)

		if len(levels) < 2 {
			log.Fatal("expected at least 2 level readings")
		}

		all_levels = append(all_levels, levels)
	}

	fmt.Println("Part 1:", part_one(all_levels))
	fmt.Println("Part 2:", part_two(all_levels))
}
