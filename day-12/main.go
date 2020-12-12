package main

import (
	"log"
	"strconv"
	"strings"
)

func main() {
	route := newRoute(data)
	ship := newShip()
	ship.followRoute(route)
	log.Println("Part One solution:", calcManhattan(0, 0, ship.x, ship.y))
}

type route []instruction

type instruction struct {
	action int
	amount int
}

func newRoute(rawRoute string) route {
	route := route{}
	mapping := strings.NewReplacer(
		"N", strconv.Itoa(north),
		"S", strconv.Itoa(south),
		"E", strconv.Itoa(east),
		"W", strconv.Itoa(west),
		"L", strconv.Itoa(turnLeft),
		"R", strconv.Itoa(turnRight),
		"F", strconv.Itoa(forward),
	)
	for _, dataLine := range strings.Split(rawRoute, "\n") {
		dataLine = mapping.Replace(dataLine)
		action, err := strconv.Atoi(dataLine[0:1])
		if err != nil {
			panic(err)
		}
		amount, err := strconv.Atoi(dataLine[1:])
		if err != nil {
			panic(err)
		}
		inst := instruction{}
		inst.action = action
		inst.amount = amount
		route = append(route, inst)
	}
	return route
}

type ship struct {
	x, y     int
	rotation int // rotation = degrees / 90
}

func newShip() *ship {
	return &ship{rotation: east}
}

var cardinalityMovement = map[int][2]int{
	north: [2]int{+0, +1},
	south: [2]int{+0, -1},
	east:  [2]int{+1, +0},
	west:  [2]int{-1, +0},
}

func (s *ship) followRoute(r route) {
	for _, inst := range r {
		var movement [2]int
		switch inst.action {
		case turnLeft:
			s.rotate(inst.amount, -1)
		case turnRight:
			s.rotate(inst.amount, +1)
		case forward:
			movement = cardinalityMovement[s.rotation]
		default:
			movement = cardinalityMovement[inst.action]
		}
		s.x += movement[0] * inst.amount
		s.y += movement[1] * inst.amount
	}
}

func (s *ship) rotate(deg, mult int) int {
	s.rotation = mod(s.rotation+mult*(deg%360)/90, 4)
	return s.rotation
}

func calcManhattan(startX, startY, endX, endY int) int {
	diffX := abs(startX - endX)
	diffY := abs(startY - endY)
	return diffX + diffY
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

// % is not module in golang ლಠ益ಠ)ლ
func mod(val, m int) int {
	val = val % m
	if val < 0 {
		val = m + val
	}
	return val
}
