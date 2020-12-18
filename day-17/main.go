package main

import (
	"log"
	"strings"
)

func main() {
	initialState := newGrid(data)
	// Part One
	calcPartOne(initialState, 6)
}

type grid struct {
	cubes map[coords]struct{}
}

func newGrid(rawGrid string) grid {
	initialState := grid{
		cubes: make(map[coords]struct{}),
	}
	for y, dataLine := range strings.Split(rawGrid, "\n") {
		for x, char := range dataLine {
			if char == '#' {
				initialState.cubes[coords{x, y, 0}] = struct{}{}
			}
		}
	}
	return initialState
}

// relevantCoordsForNextState returns a slice of all cords that need to be checked to create the next state
func (g grid) relevantCoordsForNextState() []coords {
	coordsSet := make(map[coords]struct{})
	for coord := range g.cubes {
		for _, neighbor := range coord.neighborCoords() {
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
	for _, neighbor := range c.neighborCoords() {
		if _, active := g.cubes[neighbor]; active {
			count++
		}
	}
	return count
}

type coords struct {
	x, y, z int
}

func (c coords) neighborCoords() []coords {
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
				}
				index++
			}
		}
	}
	return neighborCoords
}

/******************************************************************/
/************************* PART ONE STUFF *************************/
/******************************************************************/

func calcPartOne(initialState grid, cycles int) {
	state := initialState.nextState()
	iteration := 1 // we have already done an iteration
	for iteration < cycles {
		state = state.nextState()
		iteration++
	}
	log.Println("Solution part one:", len(state.cubes))
}

func (g grid) nextState() (newState grid) {
	newState = newGrid("")
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
