package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	maskRgx := regexp.MustCompile("^mask = (.+)$")
	memRgx := regexp.MustCompile("^mem\\[(\\d+)\\] = (\\d+)$")
	memory := make(map[uint64]uint36)

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
}

type uint36 uint64

func toInt36(i uint64) uint36 {
	mask := uint64(137438953471) // 0000000000000000000000000001111111111111111111111111111111111111
	i &= mask
	return uint36(i)
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
