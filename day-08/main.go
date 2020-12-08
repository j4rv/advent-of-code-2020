package main

import (
	"errors"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	program := newProgram(data)
	// Part One
	program.run()
	log.Println("Part One solution:", program.accumulator)
	// Part Two
	program.fixInfiniteLoop()
	log.Println("Part Two solution:", program.accumulator)
}

var errLoopDetected error = errors.New("loop detected")

type instruction struct {
	timesExec int // counter of how many times the instruction has been executed
	operation string
	argument  int
}

// swapJmpNop returns true if a swap was done (the instruction's op was either a nop or a jmp)
func (inst *instruction) swapJmpNop() bool {
	switch inst.operation {
	case opJmp:
		inst.operation = opNop
	case opNop:
		inst.operation = opJmp
	default:
		return false
	}
	return true
}

type program struct {
	nextInstIndex int
	instructions  []*instruction
	accumulator   int
}

func newProgram(rawProgram string) *program {
	program := program{}
	r := regexp.MustCompile("(\\w{3}) ([+-]\\d+)")
	for _, dataLine := range strings.Split(rawProgram, "\n") {
		match := r.FindStringSubmatch(dataLine)
		if len(match) != 3 {
			panic("Data line does not match regex: " + dataLine)
		}
		arg, err := strconv.Atoi(match[2])
		if err != nil {
			panic(err) // should not happen because of the regexp match
		}
		i := instruction{operation: match[1], argument: arg}
		program.instructions = append(program.instructions, &i)
	}
	return &program
}

func (p *program) execNext() {
	inst := p.instructions[p.nextInstIndex]
	inst.timesExec++

	switch inst.operation {
	case opNop:
		p.nextInstIndex++
	case opJmp:
		p.nextInstIndex += inst.argument
	case opAcc:
		p.accumulator += inst.argument
		p.nextInstIndex++
	default:
		panic("unsupported operation: " + inst.operation)
	}
}

// run returns errLoopDetected if an instruction was going to be executed twice
func (p *program) run() error {
	for {
		if len(p.instructions) == p.nextInstIndex {
			return nil // there was no loop!
		}

		nextInst := p.instructions[p.nextInstIndex]
		if nextInst.timesExec > 0 {
			return errLoopDetected
		}
		p.execNext()
	}
}

// brute force lol
func (p *program) fixInfiniteLoop() {
	var lastSwapIndex int
	for {
		// undo last swap attempt
		if lastSwapIndex != 0 {
			p.instructions[lastSwapIndex].swapJmpNop()
		}
		lastSwapIndex++
		// do next swap attempt
		for i := lastSwapIndex; i < len(p.instructions); i++ {
			if p.instructions[i].swapJmpNop() {
				lastSwapIndex = i
				break
			}
		}
		// test the last swap
		p.reset()
		if p.run() != errLoopDetected {
			break
		}
	}
	log.Println("Loop fixed! Index of the instruction that was swapped:", lastSwapIndex)
}

func (p *program) reset() {
	p.accumulator = 0
	for _, inst := range p.instructions {
		inst.timesExec = 0
	}
	p.nextInstIndex = 0
}
