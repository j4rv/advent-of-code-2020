package main

import (
	"log"
	"math/big"
	"strconv"
	"strings"
)

func main() {
	// Part One
	var busIDs []int64
	splitData := strings.Split(exampleData, "\n")
	startTime, err := strconv.Atoi(splitData[0])
	if err != nil {
		panic("start time not a number")
	}
	for _, rawBusID := range strings.Split(splitData[1], ",") {
		busID, err := strconv.Atoi(rawBusID)
		if err != nil {
			continue
		}
		busIDs = append(busIDs, int64(busID))
	}
	calcPartOne(int64(startTime), busIDs)

	// PartTwo
	var remainders []int64
	for i, rawBusID := range strings.Split(splitData[1], ",") {
		busID, err := strconv.Atoi(rawBusID)
		if err != nil {
			continue
		}
		remainders = append(remainders, int64((busID-i)%busID)) // Positive modulo of (-i % busID)
	}
	log.Println(busIDs, remainders)
	calcPartTwo(busIDs, remainders)
}

func calcPartOne(startTime int64, busIDs []int64) {
	for time := startTime; ; time++ {
		for _, busID := range busIDs {
			if time%busID == 0 {
				waitTime := time - startTime
				log.Println("Part One solution:", waitTime*busID)
				return
			}
		}
	}
}

func calcPartTwo(busIDs []int64, rem []int64) {
	p := big.NewInt(1) // product of all bus IDs
	for _, busID := range busIDs {
		p.Mul(p, big.NewInt(busID))
	}

	s := big.NewInt(0) // ∑(a*pa*c)
	for i := 1; i < len(busIDs); i++ {
		b, r := big.NewInt(busIDs[i]), big.NewInt(rem[i])

		pb := big.NewInt(0)
		pb.Div(p, b)

		inv := modInv(pb, b)

		m := big.NewInt(0)
		m.Mul(r, pb)
		m.Mul(m, inv)
		s.Add(s, m)
	}

	// sanity check
	for i := 0; i < len(busIDs); i++ {
		b, r := big.NewInt(busIDs[i]), big.NewInt(rem[i])
		x := big.NewInt(0)
		if x.Mod(s, b).Cmp(r) != 0 {
			log.Fatal()
		}
		log.Println(s, b, x, r)
	}

	s.Mod(s, p)
	log.Println("Part Two solution:", s)
}

// modulo inverse
func modInv(a, m *big.Int) *big.Int {
	a2 := big.NewInt(0)
	a2.Mod(a, m)
	one := big.NewInt(1)
	for i := big.NewInt(1); i.Cmp(m) < 0; i.Add(i, one) {
		ai := big.NewInt(0)
		ai.Mul(a2, i)
		if ai.Mod(ai, m).Cmp(one) == 0 {
			return i
		}
	}
	panic("mod inv not found")
}
