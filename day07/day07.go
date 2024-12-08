package main

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/agelito/go_aoc_2024/utils"
)

type Op int

const (
	OpAdd Op = iota
	OpMul
	OpConcat
)

type Calibration struct {
	result int
	values []int
}

type Expr struct {
	op Op
	l  int
	r  int
}

func evaluateExpression(expr Expr) int {
	if expr.op == OpAdd {
		return expr.l + expr.r
	} else if expr.op == OpMul {
		return expr.l * expr.r
	} else if expr.op == OpConcat {
		combined := fmt.Sprintf("%v%v", expr.l, expr.r)
		num, _ := strconv.Atoi(combined)
		return num
	}

	panic("invalid operator")
}

func evaluateExpressions(left int, rest []int, ops []Op) []int {
	results := make([]int, 0)

	for _, op := range ops {
		expr := Expr{op, left, rest[0]}
		results = append(results, evaluateExpression(expr))
	}

	if len(rest) == 1 {
		return results
	}

	inner := make([]int, 0)
	for _, result := range results {
		inner = append(inner, evaluateExpressions(result, rest[1:], ops)...)
	}

	return inner
}

func parseCalibrations(lines []string) []Calibration {
	calibrations := make([]Calibration, 0)

	for _, line := range lines {
		parts := strings.Split(line, ":")
		result, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Panic("could not parse number", err)
		}

		inValues := strings.Split(strings.Trim(parts[1], " "), " ")

		values := make([]int, 0)
		for _, inValue := range inValues {
			value, err := strconv.Atoi(inValue)
			if err != nil {
				log.Panic("could not parse number", err)
			}

			values = append(values, value)
		}

		calibrations = append(calibrations, Calibration{result, values})
	}

	return calibrations
}

func part1(lines []string) int {
	ops := []Op{OpAdd, OpMul}
	calibrations := parseCalibrations(lines)

	result := 0
	for _, calibration := range calibrations {
		results := evaluateExpressions(calibration.values[0], calibration.values[1:], ops)

		if slices.Contains(results, calibration.result) {
			result += calibration.result
		}
	}

	return result
}

func part2(lines []string) int {
	ops := []Op{OpAdd, OpMul, OpConcat}
	calibrations := parseCalibrations(lines)

	result := 0
	for _, calibration := range calibrations {
		results := evaluateExpressions(calibration.values[0], calibration.values[1:], ops)

		if slices.Contains(results, calibration.result) {
			result += calibration.result
		}
	}

	return result
}

func main() {
	lines := utils.ReadLines("input.txt")

	result1 := part1(lines)
	result2 := part2(lines)

	fmt.Println("Part 1:", result1)
	fmt.Println("Part 2:", result2)
}
