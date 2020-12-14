package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	runPartOne(data)
	runPartTwo(data)
}

func mustAtoui(s string) uint64 {
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}

// Part One

func runPartOne(data string) {
	memory := make(map[uint64]uint64)
	maskRgx := regexp.MustCompile("^mask = (.+)$")
	memRgx := regexp.MustCompile("^mem\\[(\\d+)\\] = (\\d+)$")

	var setMask, clrMask uint64
	for _, dataLine := range strings.Split(data, "\n") {
		maskMatch := maskRgx.FindStringSubmatch(dataLine)
		if len(maskMatch) != 0 {
			setMask, clrMask = processMask(maskMatch[1])
			continue
		}

		memMatch := memRgx.FindStringSubmatch(dataLine)
		address := mustAtoui(memMatch[1])
		value := mustAtoui(memMatch[2])
		value |= setMask
		value &= clrMask
		memory[address] = value
	}

	var sum uint64
	for _, val := range memory {
		sum += uint64(val)
	}
	log.Println("Part One solution:", sum)
}

func processMask(s string) (setMask uint64, clrMask uint64) {
	for i := range s {
		c := s[len(s)-1-i]
		switch c {
		case '1':
			setMask |= (1 << i)
		case '0':
			clrMask |= (1 << i)
		}
	}
	return setMask, ^clrMask
}

// Part Two

func runPartTwo(data string) {
	memory := make(map[uint64]uint64) // reset memory from part one
	maskRgx := regexp.MustCompile("^mask = (.+)$")
	memRgx := regexp.MustCompile("^mem\\[(\\d+)\\] = (\\d+)$")

	var mask string
	for _, dataLine := range strings.Split(data, "\n") {
		maskMatch := maskRgx.FindStringSubmatch(dataLine)
		if len(maskMatch) != 0 {
			mask = maskMatch[1]
			continue
		}

		memMatch := memRgx.FindStringSubmatch(dataLine)
		address := mustAtoui(memMatch[1])
		addresses := []uint64{0}
		// calculate all possible addresses by applying the mask
		// mask[i] is the the most significant bit
		// mask[35] is be the least significant bit
		// so we traverse mask from left to right and invert the bit index
		for i := 0; i < 36; i++ {
			bitIndex := 35 - i
			switch mask[i] {
			case '0':
				for j := range addresses {
					if bitAt(address, bitIndex) == 1 {
						addresses[j] = setBit(bitIndex, addresses[j])
					}
				}
			case '1':
				for j := range addresses {
					addresses[j] = setBit(bitIndex, addresses[j])
				}
			case 'X':
				for j := range addresses {
					addresses = append(addresses, addresses[j])   // with bit 0
					addresses[j] = setBit(bitIndex, addresses[j]) // with bit 1
				}
			}
		}

		value := mustAtoui(memMatch[2])
		for _, address := range addresses {
			memory[address] = value
		}
	}

	var sum uint64
	for _, val := range memory {
		sum += uint64(val)
	}
	log.Println("Part Two solution:", sum)
}

func bitAt(val uint64, index int) uint64 {
	mask := uint64(1) << index
	return (val & mask) >> index
}

func setBit(index int, b uint64) uint64 {
	return b | (1 << index)
}

func clrBit(index int, b uint64) uint64 {
	return b & ^(1 << index)
}
