package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

func part_one(a []int, b []int) int {
	total := 0

	for i := 0; i < len(a); i++ {
		total += abs_diff(a[i], b[i])
	}

	return total
}

func part_two(a []int, b []int) int {
	lookup := make(map[int]int)

	for _, num := range a {
		n, ok := lookup[num]

		if ok {
			continue
		}

		for _, bnum := range b {
			if bnum == num {
				n += 1
			}
		}

		lookup[num] = n
	}

	total := 0
	for i := 0; i < len(a); i++ {
		total += a[i] * lookup[a[i]]
	}

	return total
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	a := make([]int, 0)
	b := make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "   ")

		a1, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatal(err)
		}

		b1, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}

		a = append(a, a1)
		b = append(b, b1)
	}

	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})

	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})

	if len(a) != len(b) {
		log.Fatal("expected both lists to be of same length")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1: ", part_one(a, b))
	fmt.Println("Part 2: ", part_two(a, b))
}
