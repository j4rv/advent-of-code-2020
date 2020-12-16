package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fields, yourTicket, nearbyTickets := parse(data)

	var solution1 int
	var validTickets []ticket // for part two
	for _, nearbyTicket := range nearbyTickets {
		validTicket := true
		for _, fieldVal := range nearbyTicket {
			validFieldVal := false
			for _, field := range fields {
				if validateField(fieldVal, field) {
					validFieldVal = true
					break
				}
			}
			if !validFieldVal {
				solution1 += fieldVal
				validTicket = false
			}
		}
		if validTicket {
			validTickets = append(validTickets, nearbyTicket)
		}
	}
	log.Println("Part One solution:", solution1)

	/* Part Two Strategy:
	 * For every field, discard all possible columns by checking ticket field values
	 * Then, for every field, check if all column indexes except one has been discarded. When that happens:
	 *   Find the missing column index in the discarded set
	 *   Set it in nameToColumnIndex
	 *   Discard that index for the rest of the fields
	 *   Keep searching until we have found all fields's column index
	 * Traverse nameToColumnIndex to find the second solution
	 */

	nameToColumnIndex := make(map[string]int)
	nameToDiscardedCols := make(map[string]map[int]bool) // name to "set" of discarded column indexes

	for _, field := range fields {
		nameToDiscardedCols[field.name] = make(map[int]bool)
		for _, validTicket := range validTickets {
			for i, fieldVal := range validTicket {
				if nameToDiscardedCols[field.name][i] {
					continue
				}
				if !validateField(fieldVal, field) {
					nameToDiscardedCols[field.name][i] = true
				}
			}
		}
	}

	for len(nameToColumnIndex) < len(fields) {
		for _, field := range fields {
			// don't do anything if we have found the index for this field
			if _, ok := nameToColumnIndex[field.name]; ok {
				continue
			}

			discarded := nameToDiscardedCols[field.name]
			if len(discarded) == len(fields)-1 {
				// only one column possibility left, jackpot!
				colIndex := 0
				for i := 0; i < len(fields); i++ {
					if !discarded[i] {
						colIndex = i
						nameToColumnIndex[field.name] = colIndex
						break
					}
				}
				// discard this field's column for the other fields
				for _, f := range fields {
					if f.name == field.name {
						continue
					}
					nameToDiscardedCols[f.name][colIndex] = true
				}
			}
		}
	}

	solution2 := 1
	for name, columnIndex := range nameToColumnIndex {
		if strings.Contains(name, "departure") {
			solution2 *= yourTicket[columnIndex]
		}
	}
	log.Println("Part Two solution:", solution2)
}

var rulesRgx = regexp.MustCompile("^([a-z ]+): (\\d+)-(\\d+) or (\\d+)-(\\d+)$")

func parse(rawData string) ([]fieldRules, ticket, []ticket) {
	split := strings.Split(rawData, "\n\n")
	rawRules, rawYourTicket, rawTickets := split[0], split[1], split[2]

	var rules []fieldRules
	for _, rawRule := range strings.Split(rawRules, "\n") {
		match := rulesRgx.FindStringSubmatch(rawRule)
		if len(match) != 6 {
			panic("raw rule did not match:" + rawRule)
		}
		var rangeLo, rangeHi [2]int
		rangeLo[0], rangeLo[1] = mustAtoi(match[2]), mustAtoi(match[3])
		rangeHi[0], rangeHi[1] = mustAtoi(match[4]), mustAtoi(match[5])
		rules = append(rules, fieldRules{match[1], rangeLo, rangeHi})
	}

	rawYourTicket = strings.Replace(rawYourTicket, "your ticket:\n", "", 1)
	yourTicket := newTicket(rawYourTicket)

	var nearbyTickets []ticket
	rawTickets = strings.Replace(rawTickets, "nearby tickets:\n", "", 1)
	for _, rawTicket := range strings.Split(rawTickets, "\n") {
		nearbyTickets = append(nearbyTickets, newTicket(rawTicket))
	}

	return rules, yourTicket, nearbyTickets
}

type ticket []int

func newTicket(rawTicket string) ticket {
	var res ticket
	for _, rawNum := range strings.Split(rawTicket, ",") {
		res = append(res, mustAtoi(rawNum))
	}
	return res
}

type fieldRules struct {
	name    string
	rangeLo [2]int
	rangeHi [2]int
}

func validateField(value int, rules fieldRules) bool {
	if rules.rangeLo[0] <= value && value <= rules.rangeLo[1] {
		return true
	}
	if rules.rangeHi[0] <= value && value <= rules.rangeHi[1] {
		return true
	}
	return false
}

// I should make an aocutils package...

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
