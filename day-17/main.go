package main

import (
	"log"
	"strings"
)

func main() {
	initialState := newGrid(data)
	log.Println("Part One solution:", len(simulateCycles(initialState, 6).cubes))
	initialState.useFourthDimension = true
	log.Println("Part Two solution:", len(simulateCycles(initialState, 6).cubes))
}

func simulateCycles(initialState grid, cycles int) grid {
	state := initialState
	iteration := 0
	for iteration < cycles {
		state = state.nextState()
		iteration++
	}
	return state
}

// Coords type and methods

type coords struct {
	x, y, z, w int
}

func (c coords) neighborCoords3D() []coords {
	neighborCoords := make([]coords, 3*3*3-1)
	index := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			for k := -1; k <= 1; k++ {
				if i == 0 && j == 0 && k == 0 {
					continue
				}
				neighborCoords[index] = coords{
					c.x + i,
					c.y + j,
					c.z + k,
					0,
				}
				index++
			}
		}
	}
	return neighborCoords
}

func (c coords) neighborCoords4D() []coords {
	neighborCoords := make([]coords, 3*3*3*3-1)
	index := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			for k := -1; k <= 1; k++ {
				for l := -1; l <= 1; l++ {
					if i == 0 && j == 0 && k == 0 && l == 0 {
						continue
					}
					neighborCoords[index] = coords{
						c.x + i,
						c.y + j,
						c.z + k,
						c.w + l,
					}
					index++
				}
			}
		}
	}
	return neighborCoords
}

// Grid type and methods

type grid struct {
	cubes              map[coords]struct{}
	useFourthDimension bool
}

func newGrid(rawGrid string) grid {
	initialState := grid{
		cubes: make(map[coords]struct{}),
	}
	for y, dataLine := range strings.Split(rawGrid, "\n") {
		for x, char := range dataLine {
			if char == '#' {
				initialState.cubes[coords{x: x, y: y}] = struct{}{}
			}
		}
	}
	return initialState
}

// relevantCoordsForNextState returns a slice of all cords that need to be checked to create the next state
func (g grid) relevantCoordsForNextState() []coords {
	coordsSet := make(map[coords]struct{}, len(g.cubes))
	for coord := range g.cubes {
		var neighbors []coords
		if g.useFourthDimension {
			neighbors = coord.neighborCoords4D()
		} else {
			neighbors = coord.neighborCoords3D()
		}
		for _, neighbor := range neighbors {
			coordsSet[neighbor] = struct{}{}
		}
	}
	// map keySet to slice
	coordsSlice := make([]coords, len(coordsSet))
	i := 0
	for coord := range coordsSet {
		coordsSlice[i] = coord
		i++
	}
	return coordsSlice
}

func (g grid) activeNeighbors(c coords) int {
	var count int
	var neighbors []coords

	if g.useFourthDimension {
		neighbors = c.neighborCoords4D()
	} else {
		neighbors = c.neighborCoords3D()
	}
	for _, neighbor := range neighbors {
		if _, active := g.cubes[neighbor]; active {
			count++
		}
	}

	return count
}

func (g grid) nextState() (newState grid) {
	newState = grid{
		cubes:              make(map[coords]struct{}),
		useFourthDimension: g.useFourthDimension,
	}
	for _, coord := range g.relevantCoordsForNextState() {
		_, active := g.cubes[coord]
		activeNeighbors := g.activeNeighbors(coord)
		if active {
			if activeNeighbors == 2 || activeNeighbors == 3 {
				newState.cubes[coord] = struct{}{}
			} // else the cube becomes inactive
		} else {
			if activeNeighbors == 3 {
				newState.cubes[coord] = struct{}{}
			} // else the cube remains inactive
		}
	}
	return newState
}
