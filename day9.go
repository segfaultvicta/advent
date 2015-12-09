package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Node struct {
	Name  string
	Paths map[string]int
}

func makePath(nodes map[string]Node, from string, to string, dist int) {
	fromNode, fromExists := nodes[from]
	if fromExists {
		_, toExists := fromNode.Paths[to]
		if !toExists {
			fromNode.Paths[to] = dist
		}
	} else {
		empty := make(map[string]int)
		empty[to] = dist
		nodes[from] = Node{Name: from, Paths: empty}
	}
}

func (n *Node) bestPath(remaining map[string]Node, shortest bool) (distance int, path string) {
	// what the actual balls is happening here?
	if len(remaining) == 1 {
		for _, remainingNode := range remaining {
			return n.Paths[remainingNode.Name], n.Name + " -> " + remainingNode.Name
		}
	}

	var best int
	if shortest {
		best = 999999
	} else {
		best = 0
	}

	var bestPath string
	var bestNodeName string

	newRemaining := cut(n.Name, remaining)

	for _, node := range newRemaining {
		d, p := node.bestPath(newRemaining, shortest)
		if shortest {
			if d < best {
				best = d
				bestPath = p
				bestNodeName = node.Name
			}
		} else {
			if d > best {
				best = d
				bestPath = p
				bestNodeName = node.Name
			}
		}
	}
	return n.Paths[bestNodeName] + best, n.Name + " -> " + bestPath
}

func day9sideA(lines []string) string {
	var nodes map[string]Node
	nodes = make(map[string]Node)

	for _, line := range lines {
		split := strings.Split(line, " ")
		from := split[0]
		to := split[2]
		dist, _ := strconv.Atoi(split[4])
		makePath(nodes, from, to, dist)
		makePath(nodes, to, from, dist)
	}

	var shortestPaths []int

	for name := range nodes {
		node := nodes[name]
		remaining := cut(node.Name, nodes)
		distance, path := node.bestPath(remaining, true)
		fmt.Println("shortest path from", nodes[name].Name, "traversing all nodes is", path, "with distance", distance)
		shortestPaths = append(shortestPaths, distance)
	}
	sort.Ints(shortestPaths)
	return strconv.Itoa(shortestPaths[0])
}

func cut(nodeName string, nodes map[string]Node) map[string]Node {
	temp := make(map[string]Node)
	for _, node := range nodes {
		if node.Name != nodeName {
			temp[node.Name] = node
		}
	}
	return temp
}

func day9sideB(lines []string) string {
	var nodes map[string]Node
	nodes = make(map[string]Node)

	for _, line := range lines {
		split := strings.Split(line, " ")
		from := split[0]
		to := split[2]
		dist, _ := strconv.Atoi(split[4])
		makePath(nodes, from, to, dist)
		makePath(nodes, to, from, dist)
	}

	var longestPaths []int

	for name := range nodes {
		node := nodes[name]
		remaining := cut(node.Name, nodes)
		distance, path := node.bestPath(remaining, false)
		fmt.Println("shortest path from", nodes[name].Name, "traversing all nodes is", path, "with distance", distance)
		longestPaths = append(longestPaths, distance)
	}
	sort.Ints(longestPaths)
	return strconv.Itoa(longestPaths[len(longestPaths)-1])
}
