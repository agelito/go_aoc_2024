package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func check_dependency_order(update []int, dependencies map[int][]int) bool {
	for i := 0; i < len(update); i++ {
		for j := 0; j < i; j++ {
			dependency := update[i]
			dependent := update[j]

			dependent_pages := dependencies[dependent]

			if slices.Contains(dependent_pages, dependency) {
				return false
			}
		}
	}

	return true
}

func sort_update(update []int, dependencies map[int][]int) {
	for i := 0; i < len(update); i++ {
		for j := 0; j < i; j++ {
			dependency := update[i]
			dependent := update[j]

			dependent_pages := dependencies[dependent]

			if slices.Contains(dependent_pages, dependency) {
				update[i], update[j] = update[j], update[i]
			}
		}
	}

	if !check_dependency_order(update, dependencies) {
		log.Panic("expected update to be valid after sorting")
	}
}

func part1(updates [][]int, page_dependencies map[int][]int) int {
	result := 0
	for _, update := range updates {
		valid := check_dependency_order(update, page_dependencies)

		if valid {
			result += update[len(update)/2]
		}
	}

	return result
}

func part2(updates [][]int, page_dependencies map[int][]int) int {
	result := 0
	incorrectUpdates := make([][]int, 0)

	for _, update := range updates {
		if !check_dependency_order(update, page_dependencies) {
			incorrectUpdates = append(incorrectUpdates, update)
		}
	}

	for _, update := range incorrectUpdates {
		sort_update(update, page_dependencies)
		result += update[len(update)/2]
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

	scan_ordering := true
	ordering_inputs := make([]string, 0)
	update_inputs := make([]string, 0)

	for scanner.Scan() {
		input := scanner.Text()

		if scan_ordering && len(input) == 0 {
			scan_ordering = false
			continue
		}

		if scan_ordering {
			ordering_inputs = append(ordering_inputs, input)
		} else {
			update_inputs = append(update_inputs, input)
		}
	}

	page_dependencies := make(map[int][]int)

	for _, order := range ordering_inputs {
		ordering_pair := strings.Split(order, "|")

		dependency, _ := strconv.Atoi(ordering_pair[0])
		page_number, _ := strconv.Atoi(ordering_pair[1])

		dependencies := page_dependencies[page_number]
		dependencies = append(dependencies, dependency)

		page_dependencies[page_number] = dependencies

	}

	updates := make([][]int, 0)

	for _, update_line := range update_inputs {
		update := make([]int, 0)

		for _, page := range strings.Split(update_line, ",") {
			page_number, _ := strconv.Atoi(page)
			update = append(update, page_number)
		}

		updates = append(updates, update)
	}

	part1 := part1(updates, page_dependencies)
	part2 := part2(updates, page_dependencies)

	fmt.Println("Part 1", part1)
	fmt.Println("Part 2", part2)
}
