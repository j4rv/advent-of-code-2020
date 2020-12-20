package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	tiles := make(map[int]tile)
	for _, rawTile := range strings.Split(exampleData, "\n\n") {
		tile := tile{sprite: make([][]rune, spriteSize)}
		lines := strings.Split(rawTile, "\n")
		fmt.Sscanf(lines[0], "Tile %d:", &tile.id)
		lines = lines[1:]
		for row, line := range lines {
			for _, c := range line {
				tile.sprite[row] = append(tile.sprite[row], c)
			}
		}
		tiles[tile.id] = tile
	}
	log.Println("\n" + tiles[2311].String())
	log.Println("\n" + tiles[2311].rotateClockWise().String())
}

type tile struct {
	id     int
	sprite [][]rune // [row][col]
}

func (t tile) String() string {
	var s string
	for _, row := range t.sprite {
		s += string(row) + "\n"
	}
	return s
}

func (t tile) border(side byte) string {
	switch side {
	case up:
		return string(t.sprite[0])
	case down:
		return string(t.sprite[spriteSize-1])
	case left, right:
		var col int
		if side == left {
			col = 0
		} else {
			col = spriteSize - 1
		}
		var res string
		for i := 0; i < spriteSize; i++ {
			res += string(t.sprite[i][col])
		}
		return res
	default:
		log.Fatal("non valid side:", side)
		return ""
	}
}

func (t tile) rotateClockWise() tile {
	res := tile{sprite: make([][]rune, spriteSize)}
	for col := 0; col < spriteSize; col++ {
		for row := spriteSize - 1; row >= 0; row-- {
			res.sprite[col] = append(res.sprite[col], t.sprite[row][col])
		}
	}
	return res
}
