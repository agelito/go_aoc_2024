package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"slices"
)

type Dir int
type Visited []bool

const (
	UP Dir = iota
	RIGHT
	DOWN
	LEFT
)

type Loc struct {
	x, y int
	dir  Dir
}

type Board struct {
	tiles []bool
	w     int
	h     int
}

func turnLoc(loc Loc) Loc {
	if loc.dir == UP {
		return Loc{loc.x, loc.y, RIGHT}
	} else if loc.dir == RIGHT {
		return Loc{loc.x, loc.y, DOWN}
	} else if loc.dir == DOWN {
		return Loc{loc.x, loc.y, LEFT}
	} else if loc.dir == LEFT {
		return Loc{loc.x, loc.y, UP}
	}

	panic("invalid direction")
}

func stepLoc(loc Loc) Loc {
	if loc.dir == UP {
		return Loc{loc.x, loc.y - 1, loc.dir}
	} else if loc.dir == RIGHT {
		return Loc{loc.x + 1, loc.y, loc.dir}
	} else if loc.dir == DOWN {
		return Loc{loc.x, loc.y + 1, loc.dir}
	} else if loc.dir == LEFT {
		return Loc{loc.x - 1, loc.y, loc.dir}
	}

	panic("invalid direction")
}

func isInside(loc Loc, board Board) bool {
	return loc.x >= 0 && loc.x < board.w && loc.y >= 0 && loc.y < board.h
}

func walkGuad(loc Loc, board Board) ([][]Dir, error) {
	visited := make([][]Dir, board.w*board.h)

	for {
		index := loc.x + loc.y*board.w
		if slices.Contains(visited[index], loc.dir) {
			return visited, errors.New("GuardLoop")
		}
		visited[index] = append(visited[index], loc.dir)

		nextLoc := stepLoc(loc)

		if !isInside(nextLoc, board) {
			break
		}

		for board.tiles[nextLoc.x+nextLoc.y*board.w] {
			loc = turnLoc(loc)
			nextLoc = stepLoc(loc)
		}

		loc = nextLoc
	}

	return visited, nil
}

func createBoard(lines []string) Board {
	rows := len(lines[0])
	cols := len(lines)

	tiles := make([]bool, rows*cols)

	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			if lines[y][x] == '#' {
				tiles[x+y*cols] = true
			}
		}
	}

	board := Board{tiles, cols, rows}

	return board
}

func createLocation(lines []string) Loc {
	rows := len(lines[0])
	cols := len(lines)

	loc := Loc{0, 0, UP}

	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			if lines[y][x] == '^' {
				loc = Loc{x, y, UP}
				break
			}
		}
	}

	return loc
}

func part1(loc Loc, board Board) int {
	visited, _ := walkGuad(loc, board)

	result := 0

	for _, t := range visited {
		if len(t) > 0 {
			result += 1
		}
	}

	return result
}

func part2(loc Loc, board Board) int {
	result := 0

	startX, startY := loc.x, loc.y

	for y := 0; y < board.h; y++ {
		for x := 0; x < board.w; x++ {
			if x == startX && y == startY {
				continue
			}

			modifiedTiles := make([]bool, len(board.tiles))
			copy(modifiedTiles, board.tiles)

			modifiedTiles[x+y*board.w] = true

			_, err := walkGuad(loc, Board{modifiedTiles, board.w, board.h})

			if err != nil {
				result += 1
			}

		}
	}

	return result
}

func printBoard(board Board, visited []bool) {
	for y := 0; y < board.h; y++ {
		for x := 0; x < board.w; x++ {
			if board.tiles[x+y*board.w] {
				fmt.Print("#")
			} else if visited[x+y*board.w] {
				fmt.Print("X")
			} else {
				fmt.Print(".")
			}
		}

		fmt.Println()
	}
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

	board := createBoard(lines)
	loc := createLocation(lines)

	result1 := part1(loc, board)
	result2 := part2(loc, board)

	fmt.Println("Part 1:", result1)
	fmt.Println("Part 2:", result2)
}
