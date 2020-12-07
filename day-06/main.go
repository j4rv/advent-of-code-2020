package main

import (
	"log"
	"strings"
)

func main() {
	// Part One
	var answers []answersSet
	for _, rawGroupAnswers := range strings.Split(data, "\n\n") {
		ga := make(answersSet)
		rawGroupAnswers = strings.ReplaceAll(rawGroupAnswers, "\n", "")
		for _, answer := range rawGroupAnswers {
			ga[answer] = struct{}{}
		}
		answers = append(answers, ga)
	}

	var count1 int
	for _, a := range answers {
		count1 += len(a)
	}
	log.Println("Part One solution:", count1)

	// Part Two
	var planeGroupAnswers []groupAnswers
	for _, rawGroupAnswers := range strings.Split(data, "\n\n") {
		var ga groupAnswers
		rawPersonAnswers := strings.Split(rawGroupAnswers, "\n")
		for _, rpa := range rawPersonAnswers {
			answers := make(answersSet)
			for _, answer := range rpa {
				answers[answer] = struct{}{}
			}
			ga = append(ga, answers)
		}
		planeGroupAnswers = append(planeGroupAnswers, ga)
	}

	var count2 int
	for _, ga := range planeGroupAnswers {
		count2 += ga.countCommonAnswers()
	}
	log.Println("Part Two solution:", count2)
}

type answersSet map[rune]struct{} // For part one: A set of answers to ignore repetitions
type groupAnswers []answersSet    // For part two: Every person gets their own answers set

func (ga groupAnswers) countCommonAnswers() int {
	commonAnswers := make(answersSet)
	for _, answers := range ga {
		for answer := range answers {
			if ga.iscommonAnswer(answer) {
				commonAnswers[answer] = struct{}{}
			}
		}
	}
	return len(commonAnswers)
}

func (ga groupAnswers) iscommonAnswer(r rune) bool {
	for _, answers := range ga {
		_, ok := answers[r]
		if !ok {
			return false
		}
	}
	return true
}
