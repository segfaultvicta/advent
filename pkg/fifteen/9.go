package fifteen

import (
	"sort"
	"strconv"
	"strings"
)

type node struct {
	Name  string
	Paths map[string]int
}

func makePath(nodes map[string]node, from string, to string, dist int) {
	fromNode, fromExists := nodes[from]
	if fromExists {
		_, toExists := fromNode.Paths[to]
		if !toExists {
			fromNode.Paths[to] = dist
		}
	} else {
		empty := make(map[string]int)
		empty[to] = dist
		nodes[from] = node{Name: from, Paths: empty}
	}
}

// I don't actually know why this is a random walk, rather than a traversal.
// If you do know what I'm doing wrong - it might be a logic issue, or it
// might be not understanding Go properly - please let me know, I'm SUPER curious.
func (n node) mysteriousRandomWalk(remaining map[string]node, shortest bool) (distance int, path string) {
	//scan := bufio.NewReader(os.Stdin)
	//scan.ReadString('\n')
	newRemaining := cut(n.Name, remaining)
	//indent := strings.Repeat("   ", 7-len(newRemaining))
	//fmt.Println(indent, ">>>>>>>>>>>>>>>>>>>>>>>>>> finding best path starting at", n.Name)
	//children := ""
	//for _, tmp := range newRemaining {
	//	children += " " + tmp.Name
	//}

	//fmt.Println(indent, len(newRemaining), "children at this node", children)

	if len(newRemaining) == 1 {
		for _, remainingNode := range newRemaining {
			//fmt.Println(indent, "base case returning", n.Paths[remainingNode.Name], n.Name+" -> "+remainingNode.Name)
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

	for _, node := range newRemaining {
		//fmt.Println(indent, "about to call bestPath on", node.Name, "current known best is", best, bestPath, "(", bestNodeName, ")")
		d, p := node.mysteriousRandomWalk(newRemaining, shortest)
		//fmt.Println(indent, "in", n.Name, "just returned (", d, ",", p, ") from", node.Name)
		if shortest {
			if d < best {
				//fmt.Println(indent, "found a new best!")
				best = d
				bestPath = p
				bestNodeName = node.Name
			}
		} else {
			if d > best {
				//fmt.Println(indent, "found a new best!")
				best = d
				bestPath = p
				bestNodeName = node.Name
			}
		}
	}
	//fmt.Println(indent, "<<<<<<<<<<<<<<<<<<<<<<<<<< inductive case returning", n.Paths[bestNodeName]+best, n.Name+" -> "+bestPath)
	return n.Paths[bestNodeName] + best, n.Name + " -> " + bestPath
}

func day9sideA(lines []string) string {
	var nodes map[string]node
	nodes = make(map[string]node)

	for _, line := range lines {
		split := strings.Split(line, " ")
		from := split[0]
		to := split[2]
		dist, _ := strconv.Atoi(split[4])
		makePath(nodes, from, to, dist)
		makePath(nodes, to, from, dist)
	}

	//remaining := cut(nodes["Snowdin"].Name, nodes)
	//d, p := nodes["Snowdin"].mysteriousRandomWalk(remaining, true)
	//fmt.Println(d, p)

	var shortestPaths []int

	for name := range nodes {
		node := nodes[name]
		remaining := cut(node.Name, nodes)

		maxIter := 100
		best := 9999999999999
		for i := 0; i < maxIter; i++ {
			distance, _ := node.mysteriousRandomWalk(remaining, true)
			if distance < best {
				best = distance
			}
		}
		//distance, path := node.bestPath(remaining, true)
		//fmt.Println("shortest path from", nodes[name].Name, "traversing all nodes is", path, "with distance", distance)
		shortestPaths = append(shortestPaths, best)
	}
	sort.Ints(shortestPaths)
	return strconv.Itoa(shortestPaths[0])
}

func cut(nodeName string, nodes map[string]node) map[string]node {
	temp := make(map[string]node)
	for _, node := range nodes {
		if node.Name != nodeName {
			temp[node.Name] = node
		}
	}
	return temp
}

func day9sideB(lines []string) string {
	var nodes map[string]node
	nodes = make(map[string]node)

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

		maxIter := 100
		best := 0
		for i := 0; i < maxIter; i++ {
			distance, _ := node.mysteriousRandomWalk(remaining, false)
			if distance > best {
				best = distance
			}
		}
		//distance, path := node.bestPath(remaining, true)
		//fmt.Println("longest path from", nodes[name].Name, "traversing all nodes is", path, "with distance", distance)
		longestPaths = append(longestPaths, best)
	}
	sort.Ints(longestPaths)
	return strconv.Itoa(longestPaths[len(longestPaths)-1])
}
