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

	rgx := parse(rawRules)
	part1Rgx := regexp.MustCompile(rgx)

	var counter int
	for _, msg := range strings.Split(rawMessages, "\n") {
		if part1Rgx.MatchString(msg) {
			counter++
		}
	}
	log.Println(counter)
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
			ruleRgx := "("
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

func subparse(ruleIndex int, rawRules []string, visited map[int]string) string {
	tokens := strings.Fields(rawRules[ruleIndex])[1:] // ignore first token, it's the rule index

	// base cases
	if res, ok := visited[ruleIndex]; ok {
		return res
	}
	if len(tokens) == 1 && tokens[0][0] == '"' {
		return string(tokens[0][1])
	}

	// recursive case
	res := "("
	for _, tkn := range tokens {
		switch c := tkn[0]; {
		case '0' <= c && c <= '9':
			i := mustAtoi(tkn)
			res += subparse(i, rawRules, visited)
		case c == '|':
			res += "|"
		default:
			log.Fatal("non controlled case:", tkn)
		}
	}
	res += ")"

	visited[ruleIndex] = res

	return res
}

// I should make an aocutils package...

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
