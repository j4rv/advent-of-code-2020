package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	split := strings.Split(data, "\n\n")
	rawRules, rawMessages := split[0], split[1]
	part1Rgx := regexp.MustCompile(parse(rawRules))
	var counter int
	for _, msg := range strings.Split(rawMessages, "\n") {
		if part1Rgx.MatchString(msg) {
			counter++
		}
	}
	log.Println(counter)
}

func parse(rawRules string) string {
	return "^" + subparse(0, strings.Split(rawRules, "\n"), make(map[int]string)) + "$"
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
