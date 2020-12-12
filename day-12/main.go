package main

import (
	"log"
	"strconv"
	"strings"
)

func main() {
	// Part One
	route := newRoute(data)
	ship := newShip()
	ship.followRoutePartOne(route)
	log.Println("Part One solution:", manhattan(0, 0, ship.x, ship.y))

	// Part Two
	ship = newShip()
	ship.followRoutePartTwo(route)
	log.Println("Part Two solution:", manhattan(0, 0, ship.x, ship.y))
}

type route []instruction

type instruction struct {
	action rune
	amount int
}

func newRoute(rawRoute string) route {
	route := route{}
	for _, dataLine := range strings.Split(rawRoute, "\n") {
		amount, err := strconv.Atoi(dataLine[1:])
		if err != nil {
			panic(err)
		}
		inst := instruction{}
		inst.action = rune(dataLine[0])
		inst.amount = amount
		route = append(route, inst)
	}
	return route
}

type ship struct {
	x, y      int
	rotation  int
	waypointX int
	waypointY int
}

func newShip() *ship {
	return &ship{
		rotation:  90, // starts facing east
		waypointX: 10,
		waypointY: 1,
	}
}

var cardinalityMovement = map[rune][2]int{
	north: [2]int{+0, +1},
	south: [2]int{+0, -1},
	east:  [2]int{+1, +0},
	west:  [2]int{-1, +0},
}

func (s *ship) followRoutePartOne(r route) {
	for _, inst := range r {
		var movement [2]int
		switch inst.action {
		case turnRight:
			s.rotation += inst.amount
		case turnLeft:
			s.rotation -= inst.amount

		case forward:
			facingCardinality := rotationToCardinality[mod(s.rotation, 360)]
			movement = cardinalityMovement[facingCardinality]

		default:
			movement = cardinalityMovement[inst.action]
		}
		s.x += movement[0] * inst.amount
		s.y += movement[1] * inst.amount
	}
}

func (s *ship) followRoutePartTwo(r route) {
	for _, inst := range r {
		switch inst.action {
		case turnLeft: // invert the angles and use the turn right logic
			inst.amount = mod(-inst.amount, 360)
			fallthrough
		case turnRight:
			switch inst.amount {
			case 90:
				s.waypointX, s.waypointY = +s.waypointY, -s.waypointX
			case 180:
				s.waypointX, s.waypointY = -s.waypointX, -s.waypointY
			case 270:
				s.waypointX, s.waypointY = -s.waypointY, +s.waypointX
			}

		case forward:
			s.x += s.waypointX * inst.amount
			s.y += s.waypointY * inst.amount

		default:
			mov := cardinalityMovement[inst.action]
			s.waypointX += mov[0] * inst.amount
			s.waypointY += mov[1] * inst.amount
		}
	}
}

func manhattan(startX, startY, endX, endY int) int {
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
