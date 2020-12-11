package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"strings"
)

// flag vars
var renderStates bool

func main() {
	flag.BoolVar(&renderStates, "r", false, "if true, will produce a .png file for every calculated state")
	flag.Parse()

	initialState := newFerry(data)

	// Part One
	calcPartOne(initialState)

	// Part Two
	calcPartTwo(initialState)
}

type ferry struct {
	cells         map[coords]rune
	width, height int
}

type coords struct {
	x, y int
}

func newFerry(rawFerry string) *ferry {
	dataLines := strings.Split(rawFerry, "\n")
	initialState := &ferry{
		height: len(dataLines),
		width:  len(dataLines[0]), // all lines have the same length
		cells:  make(map[coords]rune, len(dataLines)*len(dataLines[0])),
	}
	for y, dataLine := range dataLines {
		for x, cell := range dataLine {
			initialState.cells[coords{x, y}] = cell
		}
	}
	return initialState
}

var runeWeight = map[rune]int{occupied: 1} // free seats, floor and out-of-bounds are counted as Zero

func (f *ferry) countAllOccupiedSeats() int {
	var count int
	for x := 0; x < f.width; x++ {
		for y := 0; y < f.height; y++ {
			count += runeWeight[f.cells[coords{x, y}]]
		}
	}
	return count
}

/******************************************************************/
/************************* PART ONE STUFF *************************/
/******************************************************************/

func calcPartOne(initialState *ferry) {
	if renderStates {
		renderState(initialState, "part_one(0).png")
	}
	state, changes := initialState.nextStateUsingAdjacents()
	iteration := 1 // we have already done an iteration
	for {
		if renderStates {
			renderState(state, "part_one("+fmt.Sprint(iteration)+").png")
		}
		state, changes = state.nextStateUsingAdjacents()
		if changes == 0 {
			log.Println("Part One solution:", state.countAllOccupiedSeats())
			break
		}
		iteration++
	}
}

func (f *ferry) nextStateUsingAdjacents() (newState *ferry, alterations int) {
	newState = &ferry{
		height: f.height,
		width:  f.width,
		cells:  make(map[coords]rune, f.width*f.height),
	}
	for x := 0; x < f.width; x++ {
		for y := 0; y < f.height; y++ {
			currCoords := coords{x, y}
			switch f.cells[currCoords] {
			case floor:
				newState.cells[currCoords] = floor
			case empty:
				if f.countAdjacentOccupiedSeats(currCoords) == 0 {
					newState.cells[currCoords] = occupied
					alterations++
				} else {
					newState.cells[currCoords] = empty
				}
			case occupied:
				if f.countAdjacentOccupiedSeats(currCoords) >= 4 {
					newState.cells[currCoords] = empty
					alterations++
				} else {
					newState.cells[currCoords] = occupied
				}
			default:
				log.Fatal("unknown cell type at", currCoords, ", ", f.cells[currCoords])
			}
		}
	}
	return newState, alterations
}

func (f *ferry) countAdjacentOccupiedSeats(c coords) int {
	var count int
	count += runeWeight[f.cells[coords{c.x - 1, c.y - 1}]]
	count += runeWeight[f.cells[coords{c.x - 1, c.y + 0}]]
	count += runeWeight[f.cells[coords{c.x - 1, c.y + 1}]]

	count += runeWeight[f.cells[coords{c.x + 0, c.y - 1}]]
	count += runeWeight[f.cells[coords{c.x + 0, c.y + 1}]]

	count += runeWeight[f.cells[coords{c.x + 1, c.y - 1}]]
	count += runeWeight[f.cells[coords{c.x + 1, c.y + 0}]]
	count += runeWeight[f.cells[coords{c.x + 1, c.y + 1}]]
	return count
}

/******************************************************************/
/************************* PART TWO STUFF *************************/
/******************************************************************/

func calcPartTwo(initialState *ferry) {
	if renderStates {
		renderState(initialState, "part_two(0).png")
	}
	state, changes := initialState.nextStateUsingVisibles()
	iteration := 1 // we have already done an iteration
	for {
		if renderStates {
			renderState(state, "part_two("+fmt.Sprint(iteration)+").png")
		}
		state, changes = state.nextStateUsingVisibles()
		if changes == 0 {
			log.Println("Part Two solution:", state.countAllOccupiedSeats())
			break
		}
		iteration++
	}
}

func (f *ferry) nextStateUsingVisibles() (newState *ferry, alterations int) {
	newState = &ferry{
		height: f.height,
		width:  f.width,
		cells:  make(map[coords]rune, f.width*f.height),
	}
	for x := 0; x < f.width; x++ {
		for y := 0; y < f.height; y++ {
			currCoords := coords{x, y}
			switch f.cells[currCoords] {
			case floor:
				newState.cells[currCoords] = floor
			case empty:
				if f.countVisibleOccupiedSeats(currCoords) == 0 {
					newState.cells[currCoords] = occupied
					alterations++
				} else {
					newState.cells[currCoords] = empty
				}
			case occupied:
				if f.countVisibleOccupiedSeats(currCoords) >= 5 {
					newState.cells[currCoords] = empty
					alterations++
				} else {
					newState.cells[currCoords] = occupied
				}
			default:
				log.Fatal("unknown cell type at", currCoords, ", ", f.cells[currCoords])
			}
		}
	}
	return newState, alterations
}

func (f *ferry) countVisibleOccupiedSeats(c coords) int {
	var count int
	count += runeWeight[f.findClosestSeat(c, -1, -1)]
	count += runeWeight[f.findClosestSeat(c, -1, +0)]
	count += runeWeight[f.findClosestSeat(c, -1, +1)]

	count += runeWeight[f.findClosestSeat(c, +0, -1)]
	count += runeWeight[f.findClosestSeat(c, +0, +1)]

	count += runeWeight[f.findClosestSeat(c, +1, -1)]
	count += runeWeight[f.findClosestSeat(c, +1, +0)]
	count += runeWeight[f.findClosestSeat(c, +1, +1)]
	return count
}

func (f *ferry) findClosestSeat(from coords, xDir int, yDir int) rune {
	if xDir == 0 && yDir == 0 {
		panic("provide a direction")
	}

	// Horizontal search
	if yDir == 0 {
		y := from.y
		for x := from.x + xDir; x >= 0 && x < f.width; x += xDir {
			cell := f.cells[coords{x, y}]
			if cell == floor {
				continue
			}
			return cell
		}
		return outOfBounds
	}

	// Vertical search
	if xDir == 0 {
		x := from.x
		for y := from.y + yDir; y >= 0 && y < f.height; y += yDir {
			cell := f.cells[coords{x, y}]
			if cell == floor {
				continue
			}
			return cell
		}
		return outOfBounds
	}

	// Diagonal search
	x := from.x + xDir
	y := from.y + yDir
	for {
		// bounds check
		if x < 0 || x >= f.width {
			return outOfBounds
		}
		if y < 0 || y >= f.height {
			return outOfBounds
		}

		cell := f.cells[coords{x, y}]
		if cell == floor {
			x += xDir
			y += yDir
			continue
		}
		return cell
	}
}

/****************************************************************/
/************************* RENDER STUFF *************************/
/****************************************************************/

// Thanks to https://yourbasic.org/golang/create-image/ for the example
func renderState(state *ferry, filename string) {
	w, h := state.width*8, state.height*8

	upLeft := image.Point{0, 0}
	lowRight := image.Point{w, h}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	for x := 0; x < state.width; x++ {
		for y := 0; y < state.height; y++ {
			cell := state.cells[coords{x, y}]
			drawCell(cell, x, y, img)
		}
	}

	f, err := os.Create(filename)
	if err != nil {
		panic("could not create file: " + filename)
	}
	png.Encode(f, img)
}

func drawCell(cell rune, x, y int, img *image.RGBA) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			xPixel, yPixel := x*8+i, y*8+j
			switch cell {
			case floor:
				img.Set(xPixel, yPixel, lightColor)
			case empty:
				if freeSeatSprite[i+j*8] == '#' {
					img.Set(xPixel, yPixel, darkColor)
				} else {
					img.Set(xPixel, yPixel, lightColor)
				}
			case occupied:
				if occupiedSeatSprite[i+j*8] == '#' {
					img.Set(xPixel, yPixel, darkColor)
				} else {
					img.Set(xPixel, yPixel, lightColor)
				}
			default:
				log.Println("draw cell args:", cell, x, y)
				panic("not a valid cell type")
			}
		}
	}
}

var lightColor = color.RGBA{0xEE, 0xEE, 0xEE, 0xFF}
var darkColor = color.RGBA{0x33, 0x33, 0x33, 0xFF}

// possible optimization: turn both into [8]byte, where bit = 1 -> dark pixel
const freeSeatSprite = `          ####   #    #  #    #  #    #  #    #  ######         `
const occupiedSeatSprite = `          ####   ######  ######  ######  ######  ######         `
