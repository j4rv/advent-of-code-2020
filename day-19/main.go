package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	split := strings.Split(data, "\n\n")
	rawRules, rawMessages := split[0], split[1]
	solvePartOne(rawRules, rawMessages)
	solvePartTwo(rawRules, rawMessages)
}

func solvePartOne(rawRules, rawMessages string) {
	rgx := regexp.MustCompile(parse(rawRules))

	var counter int
	for _, msg := range strings.Split(rawMessages, "\n") {
		if rgx.MatchString(msg) {
			counter++
		}
	}
	log.Println("Part One solution:", counter)
}

func solvePartTwo(rawRules, rawMessages string) {
	unlooped := unloopWithDepth(rawRules, 6) // at 6+ depth, the answer is still 316. It will vary depending on the input
	rgx := regexp.MustCompile(parse(unlooped))

	var counter int
	for _, msg := range strings.Split(rawMessages, "\n") {
		if rgx.MatchString(msg) {
			counter++
		}
	}
	log.Println("Part Two solution:", counter)
}

func parse(rawRules string) string {
	mem := make(map[int]string) // rule index to that rule's regexp
	rules := strings.Split(rawRules, "\n")
	for {
		if _, ok := mem[0]; ok {
			break
		}
	rulesLoop:
		for _, rule := range rules {
			tokens := strings.Fields(rule) // ignore first token, it's the rule index
			var i int
			fmt.Sscanf(tokens[0], "%d:", &i)
			// visited check:
			if _, ok := mem[i]; ok {
				continue
			}
			tokens = tokens[1:] // remove the index from the tokens
			// base case:
			if len(tokens) == 1 && tokens[0][0] == '"' {
				mem[i] = string(tokens[0][1])
				continue
			}
			// list of rules:
			ruleRgx := "(?:"
			for _, tkn := range tokens {
				switch c := tkn[0]; {
				case '0' <= c && c <= '9':
					tknN := mustAtoi(tkn)
					subRuleRgx, ok := mem[tknN]
					if !ok {
						continue rulesLoop
					}
					ruleRgx += subRuleRgx
				case c == '|':
					ruleRgx += "|"
				case c == '+':
					ruleRgx += "+"
				default:
					log.Fatal("non controlled case:", tkn)
				}
			}
			ruleRgx += ")"
			mem[i] = ruleRgx
		}
	}
	// we are interested in an exact match of rule 0
	return "^" + mem[0] + "$"
}

// for part two:
// generates x possible loop variations (loop depth) for rules with loops (8 and 11)
func unloopWithDepth(rawRules string, depth int) string {
	var unlooped string
	for _, line := range strings.Split(rawRules, "\n") {
		newLine := line
		// unloop rule 8
		if string(line[0:2]) == "8:" {
			newLine = "8: 42 +"
		}
		// unloop rule 11
		if string(line[0:3]) == "11:" {
			rule := "42 31"
			for i := 2; i < depth; i++ {
				rule += " | "
				for j := 0; j < i; j++ {
					rule += " 42 "
				}
				for j := 0; j < i; j++ {
					rule += " 31 "
				}
			}
			newLine = "11: " + rule
		}
		unlooped += newLine + "\n"
	}
	return unlooped[:len(unlooped)-1]
}

// I should make an aocutils package...

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
