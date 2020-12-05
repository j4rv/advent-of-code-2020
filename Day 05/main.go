package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Part One
	var maxSeatID int
	for _, seat := range strings.Split(data, "\n") {
		row, col := seatToBytes(seat)
		seatID := calcSeatID(row, col)
		if maxSeatID < seatID {
			maxSeatID = seatID
		}
	}
	log.Println("Part One solution:", maxSeatID)

	// Part Two
	// check from row 1 to 126 to ignore first and last row:
	//   if seatID is not in data AND seatIDs +1 and -1 are there: that's your seat!
	// probably slow (it's brute force after all lol) but can't think of anything better atm :(
	seatIDs := make(map[int]struct{})
	for _, seat := range strings.Split(data, "\n") {
		row, col := seatToBytes(seat)
		seatID := calcSeatID(row, col)
		seatIDs[seatID] = struct{}{}
	}
	for r := byte(1); r <= 126; r++ {
		for c := byte(0); c <= 7; c++ {
			seatID := calcSeatID(r, c)
			_, ok := seatIDs[seatID]
			if !ok {
				_, prevSeatIDOk := seatIDs[seatID-1]
				_, nextSeatIDOk := seatIDs[seatID+1]
				if prevSeatIDOk && nextSeatIDOk {
					log.Println("Part Two solution:", seatID)
					os.Exit(0)
				}
			}
		}
	}
}

func calcSeatID(row, col byte) int {
	return int(row)*8 + int(col)
}

func seatToBytes(s string) (row, col byte) {
	r := s[0:7]
	r = strings.ReplaceAll(r, "F", "0")
	r = strings.ReplaceAll(r, "B", "1")

	c := s[7:10]
	c = strings.ReplaceAll(c, "L", "0")
	c = strings.ReplaceAll(c, "R", "1")

	row = stringToByte(r)
	col = stringToByte(c)
	return row, col
}

func stringToByte(s string) byte {
	if len(s) > 8 {
		panic("length too long for a byte: " + s)
	}
	var res byte
	for i, c := range s {
		// string[i] -> "byte[7-i]"
		index := len(s) - 1 - i
		switch c {
		case '0':
			res = clrBit(index, res)
		case '1':
			res = setBit(index, res)
		default:
			panic("not a 0 or 1 character: " + string(c))
		}
	}
	return res
}

func setBit(index int, b byte) byte {
	if index < 0 || index > 7 {
		panic(fmt.Sprintf("index out of bounds: %d", index))
	}
	return b | (1 << index)
}

func clrBit(index int, b byte) byte {
	if index < 0 || index > 7 {
		panic(fmt.Sprintf("index out of bounds: %d", index))
	}
	return b & ^(1 << index)
}
