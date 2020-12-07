package main

import (
	"log"
	"strings"
)

func main() {
	f := newForest(data)

	// Part One
	log.Println("Part One solution:", f.checkSlopeStrategy(slopeStrategy{1, 3}))

	// Part Two
	strategies := []slopeStrategy{
		slopeStrategy{1, 1},
		slopeStrategy{1, 3},
		slopeStrategy{1, 5},
		slopeStrategy{1, 7},
		slopeStrategy{2, 1},
	}
	treeProduct := 1
	for _, strat := range strategies {
		trees := f.checkSlopeStrategy(strat)
		treeProduct *= trees
	}
	log.Println("Part Two solution:", treeProduct)
}

type forest struct {
	rows   []string
	width  int
	height int
}

type slopeStrategy struct {
	rowInclination int
	colInclination int
}

func newForest(str string) *forest {
	rows := strings.Split(str, "\n")
	return &forest{
		rows:   rows,
		width:  len(rows[0]),
		height: len(rows),
	}
}

func (f forest) getCell(row, col int) rune {
	if row < 0 || row >= f.height {
		panic("row out of bounds")
	}
	col = col % f.width
	return rune(f.rows[row][col])
}

// checkSlopeStrategy returns the amount of trees found in the forest
// for a slope that starts at (0, 0) with a certain inclination
func (f forest) checkSlopeStrategy(strat slopeStrategy) int {
	var trees int
	// start at one to skip the first cell (0, 0)
	for i := 1; i < f.height; i++ {
		row := i * strat.rowInclination
		col := i * strat.colInclination
		if row >= f.height {
			break
		}
		if f.getCell(row, col) == treeCell {
			trees++
		}
	}
	return trees
}
