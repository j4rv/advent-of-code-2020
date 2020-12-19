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

func unloopWithDepth(rawRules string, depth int) string {
	var unlooped string
	for _, line := range strings.Split(rawRules, "\n") {
		newLine := line
		// unloop rule 8
		if string(line[0:2]) == "8:" {
			rule := "42"
			for i := 2; i < depth; i++ {
				rule += " | "
				for j := 0; j < i; j++ {
					rule += " 42 "
				}
			}
			newLine = "8: " + rule
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

func parse(rawRules string) string {
	mem := make(map[int]string)
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
			tokens = tokens[1:]
			// visited check:
			if _, ok := mem[i]; ok {
				continue
			}
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
				default:
					log.Fatal("non controlled case:", tkn)
				}
			}
			ruleRgx += ")"
			mem[i] = ruleRgx
		}
	}
	return "^" + mem[0] + "$"
}

// I should make an aocutils package...

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
