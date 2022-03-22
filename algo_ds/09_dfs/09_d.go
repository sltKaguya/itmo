package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var array []int
var component int = 1

type Node struct {
	vertex int
	next   *Node
}

type LinkedList struct {
	state     int //0 for not seen, 1 for wip, 2 for done
	first     *Node
	component int
}

type Graph struct {
	adjList []LinkedList
}

func (list *LinkedList) pushNode(vertex int) {
	list.first = &Node{vertex: vertex, next: list.first}
}

func (g *Graph) dfsVisit(vertex int) {
	g.adjList[vertex].state = 1

	elem := g.adjList[vertex].first
	for elem != nil {
		if g.adjList[elem.vertex].state == 0 {
			g.dfsVisit(elem.vertex)
		}
		elem = elem.next
	}
	array = append(array, vertex)
}

func (g *Graph) dfsVisitInvert(vertex int) {
	g.adjList[vertex].component = component

	elem := g.adjList[vertex].first
	for elem != nil {
		if g.adjList[elem.vertex].component == 0 {
			g.dfsVisitInvert(elem.vertex)
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

	myGraph := &Graph{adjList: make([]LinkedList, vertices)}
	for i := range myGraph.adjList {
		myGraph.adjList[i].state = 0
		myGraph.adjList[i].first = nil
		myGraph.adjList[i].component = 0
	}

	myGraphInvert := &Graph{adjList: make([]LinkedList, vertices)}
	for i := range myGraphInvert.adjList {
		myGraphInvert.adjList[i].state = 0
		myGraphInvert.adjList[i].first = nil
		myGraphInvert.adjList[i].component = 0
	}

	for scanner.Scan() {
		line = strings.Fields(scanner.Text())
		x, _ := strconv.Atoi(line[0])
		y, _ := strconv.Atoi(line[1])
		myGraph.adjList[x-1].pushNode(y - 1)
		myGraphInvert.adjList[y-1].pushNode(x - 1)
	}

	for i := range myGraph.adjList {
		if myGraph.adjList[i].state == 0 {
			myGraph.dfsVisit(i)
		}
	}

	for i := range myGraphInvert.adjList {
		if myGraphInvert.adjList[i].component == 0 {
			myGraphInvert.dfsVisitInvert(array[len(array)-1])
			component++
		}

		array = array[:len(array)-1]
	}

	fout, _ := os.Create("cond.out")
	fmt.Fprintln(fout, component)
	for _, i := range myGraphInvert.adjList {
		fmt.Fprint(fout, i.component, " ")
	}
}
