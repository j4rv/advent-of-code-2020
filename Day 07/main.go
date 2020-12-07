package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	bagsGraph := newGraph()
	childBagsRgx := regexp.MustCompile(`(\d+) (\w+ \w+) bags?`)

	// process the data to create the graph:
	for _, dataLine := range strings.Split(data, "\n") {
		dataLine = dataLine[:len(dataLine)-1] // remove dot at the end of the line to simplify

		// in a line, left of " contain " is the parent key, the rest are its children bags
		split := strings.Split(dataLine, " contain ")
		rawParentKey, rawChildren := split[0], split[1]
		parentKey := strings.ReplaceAll(rawParentKey, " bags", "")

		// children are split by commas
		for _, rawChild := range strings.Split(rawChildren, ", ") {
			bagsGraph.addNode(parentKey)
			if rawChild == "no other bags" {
				continue
			}
			childData := childBagsRgx.FindStringSubmatch(rawChild)
			if len(childData) == 0 {
				log.Fatal("Child data did not match regex:", childData)
			}
			amount, _ := strconv.Atoi(childData[1])
			childKey := childData[2]
			bagsGraph.addContent(parentKey, childKey, amount)
		}
	}

	// Part One
	myBag := bagsGraph.keysToNodes["shiny gold"]
	log.Println("Part One solution:", len(bagsGraph.getAllAscendants(myBag)))

	// Part Two
	log.Println("Part Two solution:", bagsGraph.childrenWeight(myBag))
}

// Weighted directed graph structure for the bags containing bags rules:

type graph struct {
	keysToNodes map[string]*node
}

type edge struct {
	parent *node
	child  *node
	weight int
}

type node struct {
	key         string
	parentEdges []edge
	childEdges  []edge
}

func newGraph() *graph {
	return &graph{
		keysToNodes: make(map[string]*node),
	}
}

func (g *graph) addNode(key string) *node {
	existing, ok := g.keysToNodes[key]
	if !ok {
		node := &node{key: key}
		g.keysToNodes[key] = node
		return node
	}
	return existing
}

// parentKey: bag container key, ex: "bright white"
// childKey:  bag contained by parent, ex: "shiny gold"
// amount:    amount of child bags contained by the parent
func (g *graph) addContent(parentKey string, childKey string, amount int) {
	// get or add parent node
	parent, ok := g.keysToNodes[parentKey]
	if !ok {
		parent = g.addNode(parentKey)
	}
	// get or add child node
	child, ok := g.keysToNodes[childKey]
	if !ok {
		child = g.addNode(childKey)
	}
	// add weighted edge
	parentToChild := edge{parent: parent, child: child, weight: amount}
	parent.childEdges = append(parent.childEdges, parentToChild)
	child.parentEdges = append(child.parentEdges, parentToChild)
}

// getAllAscendants gets all nodes that are parent, grandparents, etc... of node n (recursive)
func (g *graph) getAllAscendants(n *node) []*node {
	ascendantsSet := make(map[*node]struct{})
	for _, pe := range n.parentEdges {
		ascendantsSet[pe.parent] = struct{}{}
		for _, grandpa := range g.getAllAscendants(pe.parent) {
			ascendantsSet[grandpa] = struct{}{}
		}
	}
	keys := make([]*node, 0, len(ascendantsSet))
	for k := range ascendantsSet {
		keys = append(keys, k)
	}
	return keys
}

// childrenWeight returns the sum of the weights of all children of node n (recursive)
func (g *graph) childrenWeight(n *node) int {
	var weight int
	for _, ce := range n.childEdges {
		weight += ce.weight                              // contains W bags
		weight += g.childrenWeight(ce.child) * ce.weight // weight of each child bag * number of bags
	}
	return weight
}
