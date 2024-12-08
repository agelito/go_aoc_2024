package main

import (
	"fmt"

	"github.com/agelito/go_aoc_2024/utils"
)

type FrequencyId byte

type Loc struct {
	x int
	y int
}

type World struct {
	antennas  []FrequencyId
	antinodes []bool
	w         int
	h         int
}

type AntennaLocations map[FrequencyId][]Loc

func isValidFrequency(freq byte) bool {
	return freq != '.'
}

func printWorld(world World) {
	for y := 0; y < world.h; y++ {
		for x := 0; x < world.w; x++ {
			freq := world.antennas[x+y*world.w]
			if world.antinodes[x+y*world.w] {
				fmt.Print("#")
			} else if freq != 0 {
				fmt.Print(string(freq))
			} else {
				fmt.Print(".")
			}
		}

		fmt.Println()
	}
}

func createResonantHarmonicsAntinodes(locations []Loc, w int, h int) []Loc {
	antinodeLocations := make([]Loc, 0)

	for i := 0; i < len(locations); i++ {
		for j := i + 1; j < len(locations); j++ {
			loc1 := locations[i]
			loc2 := locations[j]

			deltaX := loc1.x - loc2.x
			deltaY := loc1.y - loc2.y

			nodeA := Loc{loc1.x - deltaX, loc1.y - deltaY}
			for nodeA.x >= 0 && nodeA.x < w && nodeA.y >= 0 && nodeA.y < h {
				antinodeLocations = append(antinodeLocations, nodeA)
				nodeA = Loc{nodeA.x - deltaX, nodeA.y - deltaY}
			}

			nodeB := Loc{loc2.x + deltaX, loc2.y + deltaY}
			for nodeB.x >= 0 && nodeB.x < w && nodeB.y >= 0 && nodeB.y < h {
				antinodeLocations = append(antinodeLocations, nodeB)
				nodeB = Loc{nodeB.x + deltaX, nodeB.y + deltaY}
			}
		}
	}

	return antinodeLocations
}

func createAntinodes(locations []Loc, w int, h int) []Loc {
	antinodeLocations := make([]Loc, 0)

	for i := 0; i < len(locations); i++ {
		for j := i + 1; j < len(locations); j++ {
			loc1 := locations[i]
			loc2 := locations[j]

			dirX := loc1.x - loc2.x
			dirY := loc1.y - loc2.y

			nodeA := Loc{loc1.x + dirX, loc1.y + dirY}
			nodeB := Loc{loc2.x - dirX, loc2.y - dirY}

			if nodeA.x >= 0 && nodeA.x < w && nodeA.y >= 0 && nodeA.y < h {
				antinodeLocations = append(antinodeLocations, Loc{loc1.x + dirX, loc1.y + dirY})
			}

			if nodeB.x >= 0 && nodeB.x < w && nodeB.y >= 0 && nodeB.y < h {
				antinodeLocations = append(antinodeLocations, Loc{loc2.x - dirX, loc2.y - dirY})
			}
		}
	}

	return antinodeLocations
}

func part1(world World, antennaLocations AntennaLocations) int {
	antinodes := make([]Loc, 0)
	for _, locations := range antennaLocations {
		antinodes = append(antinodes, createAntinodes(locations, world.w, world.h)...)
	}

	for _, antinode := range antinodes {
		world.antinodes[antinode.x+antinode.y*world.w] = true
	}

	result := 0
	for _, antinode := range world.antinodes {
		if antinode {
			result += 1
		}
	}
	return result
}

func part2(world World, antennaLocations AntennaLocations) int {
	antinodes := make([]Loc, 0)
	for _, locations := range antennaLocations {
		antinodes = append(antinodes, createResonantHarmonicsAntinodes(locations, world.w, world.h)...)
	}

	for _, antinode := range antinodes {
		world.antinodes[antinode.x+antinode.y*world.w] = true
	}

	result := 0
	for _, antinode := range world.antinodes {
		if antinode {
			result += 1
		}
	}
	return result
}

func main() {
	lines := utils.ReadLines("input.txt")

	cols := len(lines[0])
	rows := len(lines)

	antennaLocations := make(AntennaLocations, 0)

	world := World{make([]FrequencyId, cols*rows), make([]bool, cols*rows), cols, rows}

	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			id := lines[y][x]

			if !isValidFrequency(id) {
				continue
			}

			freq := FrequencyId(id)

			world.antennas[x+y*cols] = freq

			location := Loc{x, y}

			antennas := antennaLocations[freq]
			antennas = append(antennas, location)

			antennaLocations[freq] = antennas
		}
	}

	result1 := part1(world, antennaLocations)
	result2 := part2(world, antennaLocations)

	fmt.Println("Part 1:", result1)
	fmt.Println("Part 2:", result2)
}
