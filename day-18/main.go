package main

import (
	"log"
	"strconv"
	"strings"
)

func main() {
	// Part One
	var sum1 int
	precedences := map[operator]int8{ADD: 1, MUL: 1}
	for _, exp := range strings.Split(data, "\n") {
		sum1 += solve(exp, precedences)
	}
	log.Println("Part One solution:", sum1)

	// Part Two
	var sum2 int
	precedences = map[operator]int8{ADD: 2, MUL: 1}
	for _, exp := range strings.Split(data, "\n") {
		sum2 += solve(exp, precedences)
	}
	log.Println("Part Two solution:", sum2)
}

func solve(exp string, precedences map[operator]int8) int {
	var vals []int
	var ops []operator

	calcTop := func() {
		op := ops[len(ops)-1]
		ops = ops[:len(ops)-1]
		val1, val2 := vals[len(vals)-2], vals[len(vals)-1]
		vals = vals[:len(vals)-2]
		res := calc(op, val1, val2)
		vals = append(vals, res)
	}

	calcWhileOps := func(stop func() bool) {
		for len(ops) != 0 {
			if stop != nil && stop() {
				break
			}
			calcTop()
		}
	}

	rep := strings.NewReplacer("(", " ( ", ")", " ) ")
	tokens := strings.Fields(rep.Replace(exp))
	for _, token := range tokens {
		switch token {
		case "*":
			calcWhileOps(func() bool {
				return cmpOperators(MUL, ops[len(ops)-1], precedences) < 0
			})
			ops = append(ops, MUL)
		case "+":
			calcWhileOps(func() bool {
				return cmpOperators(ADD, ops[len(ops)-1], precedences) < 0
			})
			ops = append(ops, ADD)
		case "(":
			ops = append(ops, LEP)
		case ")":
			calcWhileOps(func() bool {
				return ops[len(ops)-1] == LEP
			})
			ops = ops[:len(ops)-1] // pop the left parenthesis
		default:
			val, err := strconv.Atoi(token)
			if err != nil {
				panic("non op token and not a number: " + token)
			}
			vals = append(vals, val)
		}
	}
	calcWhileOps(nil)

	return vals[0]
}

func calc(op operator, a, b int) int {
	switch op {
	case ADD:
		return a + b
	case MUL:
		return a * b
	default:
		log.Fatal("op not supported:", op)
		return 0 // log fatal will panic, but we need this because Go's compilator doesn't know that ¯\_(ツ)_/¯
	}
}

func cmpOperators(a, b operator, precedences map[operator]int8) int8 {
	return precedences[b] - precedences[a]
}
