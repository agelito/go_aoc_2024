package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func extract_columns(lines []string) []string {
	row_count := len(lines)
	column_count := len(lines[0])

	columns := make([]string, column_count)

	for x := 0; x < column_count; x++ {
		line := make([]byte, row_count)

		for y := 0; y < row_count; y++ {
			line[y] = lines[y][x]
		}

		columns[x] = string(line)
	}

	return columns
}

func extract_diagonals(lines []string) []string {
	row_count := len(lines)
	column_count := len(lines[0])

	diagonals := make([]string, column_count)

	for x := 0; x < column_count; x++ {
		y := 0
		line := make([]byte, 0)
		for x1 := x; x1 < column_count; x1++ {
			line = append(line, lines[y][x1])
			y = y + 1
		}

		if len(line) > 3 {
			diagonals = append(diagonals, string(line))
		}
	}

	for y := 1; y < row_count; y++ {
		x := 0
		line := make([]byte, 0)

		for y1 := y; y1 < row_count; y1++ {
			line = append(line, lines[y1][x])
			x = x + 1
		}

		if len(line) > 3 {
			diagonals = append(diagonals, string(line))
		}
	}

	for x := 0; x < column_count; x++ {
		y := row_count - 1
		line := make([]byte, 0)

		for x1 := x; x1 < column_count; x1++ {
			line = append(line, lines[y][x1])
			y = y - 1
		}

		if len(line) > 3 {
			diagonals = append(diagonals, string(line))
		}
	}

	for y := row_count - 2; y > 0; y-- {
		x := 0
		line := make([]byte, 0)

		for y1 := y; y1 >= 0; y1-- {
			line = append(line, lines[y1][x])
			x = x + 1
		}

		if len(line) > 3 {
			diagonals = append(diagonals, string(line))
		}
	}

	return diagonals
}

func part1(lines []string) int {
	xmas := []string{"XMAS", "SAMX"}

	result := 0

	all_lines := make([]string, 0)
	all_lines = append(all_lines, lines...)
	all_lines = append(all_lines, extract_columns(lines)...)
	all_lines = append(all_lines, extract_diagonals(lines)...)

	for _, match := range xmas {
		for _, line := range all_lines {
			result += strings.Count(line, match)
		}
	}

	return result
}

type Loc struct {
	x, y int
}

func matchLoc(loc Loc, lines []string) bool {
	tl := lines[loc.y-1][loc.x-1]
	bl := lines[loc.y+1][loc.x-1]
	tr := lines[loc.y-1][loc.x+1]
	br := lines[loc.y+1][loc.x+1]

	a := tl == byte('M') && br == byte('S') || tl == byte('S') && br == byte('M')
	b := bl == byte('M') && tr == byte('S') || bl == byte('S') && tr == byte('M')

	return a && b
}

func part2(lines []string) int {
	rows := len(lines)
	cols := len(lines[0])

	search := make([]Loc, 0)

	for y := 1; y < rows-1; y++ {
		for x := 1; x < cols-1; x++ {
			if lines[y][x] == byte('A') {
				search = append(search, Loc{x, y})
			}
		}
	}

	result := 0

	for _, loc := range search {
		if matchLoc(loc, lines) {
			result += 1
		}
	}

	return result
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := make([]string, 0)

	for scanner.Scan() {
		input := scanner.Text()
		lines = append(lines, input)
	}

	result1 := part1(lines)
	result2 := part2(lines)

	fmt.Println("Part 1:", result1)
	fmt.Println("Part 2:", result2)
}
