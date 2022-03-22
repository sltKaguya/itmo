package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	vertex int
	next   *Node
}

type LinkedList struct {
	state int //0 for not seen, 1 for wip, 2 for done
	first *Node
}

type Graph struct {
	partCheck int
	adjList   []LinkedList
}

func (list *LinkedList) pushNode(vertex int) {
	list.first = &Node{vertex: vertex, next: list.first}
}

func (g *Graph) dfsVisit(vertex, state int) {
	if state == 1 {
		g.adjList[vertex].state = 2
	} else {
		g.adjList[vertex].state = 1
	}

	elem := g.adjList[vertex].first
	for elem != nil {
		if g.adjList[elem.vertex].state == 0 {
			g.dfsVisit(elem.vertex, g.adjList[vertex].state)
		} else if g.adjList[vertex].state == g.adjList[elem.vertex].state {
			g.partCheck = -1
			return
		}
		elem = elem.next
	}
}

func main() {
	fin, _ := os.Open("bipartite.in")
	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	line := strings.Fields(scanner.Text())
	vertices, _ := strconv.Atoi(line[0])

	myGraph := &Graph{partCheck: 0, adjList: make([]LinkedList, vertices)}
	for i := range myGraph.adjList {
		myGraph.adjList[i].state = 0
		myGraph.adjList[i].first = nil
	}

	for scanner.Scan() {
		line = strings.Fields(scanner.Text())
		x, _ := strconv.Atoi(line[0])
		y, _ := strconv.Atoi(line[1])
		myGraph.adjList[x-1].pushNode(y - 1)
		myGraph.adjList[y-1].pushNode(x - 1)
	}

	for i := range myGraph.adjList {
		if myGraph.adjList[i].state == 0 {
			myGraph.dfsVisit(i, 1)
		}
	}

	fout, _ := os.Create("bipartite.out")

	if myGraph.partCheck == 0 {
		fmt.Fprint(fout, "YES")
	} else {
		fmt.Fprint(fout, "NO")
	}
}
