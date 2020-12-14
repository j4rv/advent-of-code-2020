package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// Part One
	memory := make(map[uint64]uint36)
	maskRgx := regexp.MustCompile("^mask = (.+)$")
	memRgx := regexp.MustCompile("^mem\\[(\\d+)\\] = (\\d+)$")
	setMask, clrMask := uint64(0), ^uint64(0)
	for _, dataLine := range strings.Split(data, "\n") {
		maskMatch := maskRgx.FindStringSubmatch(dataLine)
		if len(maskMatch) != 0 {
			setMask, clrMask = processMask(maskMatch[1])
			continue
		}
		memMatch := memRgx.FindStringSubmatch(dataLine)
		value := mustAtoui(memMatch[2])
		value |= setMask
		value &= clrMask
		address := mustAtoui(memMatch[1])
		memory[address] = toInt36(value)
	}

	var sum uint64
	for _, val := range memory {
		sum += uint64(val)
	}
	log.Println("Part One solution:", sum)

	// Part Two
	memory = make(map[uint64]uint36) // reset memory from part one
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
			memory[address] = toInt36(value)
		}
	}

	sum = 0
	for _, val := range memory {
		sum += uint64(val)
	}
	log.Println("Part Two solution:", sum)
}

type uint36 uint64

func toInt36(i uint64) uint36 {
	mask := uint64(137438953471) // 0000000000000000000000000001111111111111111111111111111111111111
	i &= mask
	return uint36(i)
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

func processMask(s string) (uint64, uint64) {
	var setMask, clrMask uint64
	for i := range s {
		c := s[len(s)-1-i]
		switch c {
		case 'X':
			continue
		case '1':
			setMask |= (1 << i)
		case '0':
			clrMask |= (1 << i)
		}
	}
	return setMask, ^clrMask
}

func mustAtoui(s string) uint64 {
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}
