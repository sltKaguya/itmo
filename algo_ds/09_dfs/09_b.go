package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var root int = -1
var precycle []int
var cycle []int

type Node struct {
	vertex int
	next   *Node
}

type LinkedList struct {
	state int //0 for not seen, 1 for wip, 2 for done
	first *Node
}

type Graph struct {
	cycle   bool
	adjList []LinkedList
}

func (list *LinkedList) pushNode(vertex int) {
	list.first = &Node{vertex: vertex, next: list.first}
}

func (g *Graph) dfsVisit(vertex int) {
	g.adjList[vertex].state = 1
	precycle = append(precycle, vertex)
	elem := g.adjList[vertex].first

	for elem != nil {
		if g.adjList[elem.vertex].state == 0 {
			g.dfsVisit(elem.vertex)
		} else if g.adjList[elem.vertex].state == 1 && len(cycle) == 0 {
			g.cycle = true
			root = elem.vertex

			for precycle[len(precycle)-1] != root {
				cycle = append(cycle, precycle[len(precycle)-1])
				precycle = precycle[:len(precycle)-1]
			}

			cycle = append(cycle, precycle[len(precycle)-1])
			return
		}

		elem = elem.next
	}

	if len(precycle) == 0 {
		precycle = precycle[:len(precycle)-1]
	}
	g.adjList[vertex].state = 2
	return
}

func main() {
	fin, _ := os.Open("cycle.in")
	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	line := strings.Fields(scanner.Text())
	vertices, _ := strconv.Atoi(line[0])

	myGraph := &Graph{cycle: false, adjList: make([]LinkedList, vertices)}
	for i := range myGraph.adjList {
		myGraph.adjList[i].state = 0
		myGraph.adjList[i].first = nil
	}

	for scanner.Scan() {
		line = strings.Fields(scanner.Text())
		x, _ := strconv.Atoi(line[0])
		y, _ := strconv.Atoi(line[1])
		myGraph.adjList[x-1].pushNode(y - 1)
	}

	for i := range myGraph.adjList {
		if myGraph.adjList[i].state == 0 {
			myGraph.dfsVisit(i)
		}
	}

	fout, _ := os.Create("cycle.out")
	if !myGraph.cycle {
		fmt.Fprintln(fout, "NO")
	} else {
		fmt.Fprintln(fout, "YES")
		for _, i := range cycle {
			fmt.Fprint(fout, i+1, " ")
		}
	}
}
