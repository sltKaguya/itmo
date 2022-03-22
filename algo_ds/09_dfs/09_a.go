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
	state int //0 for not seen, 1 for wip, 2, for done
	first *Node
	time  int
}

type Graph struct {
	cycle   bool
	adjList []LinkedList
}

func (list *LinkedList) pushNode(vertex int) {
	list.first = &Node{vertex: vertex, next: list.first}
}

func (g *Graph) dfsVisit(vertex, time int, answer *[]int) int {
	time++
	g.adjList[vertex].state = 1
	elem := g.adjList[vertex].first

	for elem != nil {
		if g.adjList[elem.vertex].state == 0 {
			time = g.dfsVisit(elem.vertex, time, answer)
		} else if g.adjList[elem.vertex].state == 1 {
			g.cycle = true
		}
		elem = elem.next
	}
	g.adjList[vertex].state = 2
	time++
	g.adjList[vertex].time = time
	*answer = append(*answer, vertex)

	return time
}

func main() {
	fin, _ := os.Open("topsort.in")
	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	line := strings.Fields(scanner.Text())
	vertices, _ := strconv.Atoi(line[0])

	myGraph := &Graph{cycle: false, adjList: make([]LinkedList, vertices)}
	for i := range myGraph.adjList {
		myGraph.adjList[i].state = 0
		myGraph.adjList[i].time = 0
		myGraph.adjList[i].first = nil
	}

	for scanner.Scan() {
		line = strings.Fields(scanner.Text())
		x, _ := strconv.Atoi(line[0])
		y, _ := strconv.Atoi(line[1])
		myGraph.adjList[x-1].pushNode(y - 1)
	}

	time := 0
	var answer []int
	for i := range myGraph.adjList {
		if myGraph.adjList[i].state == 0 {
			time = myGraph.dfsVisit(i, time, &answer)
		}
	}

	fout, _ := os.Create("topsort.out")
	if myGraph.cycle {
		fmt.Fprintln(fout, "-1")
	} else {
		for i := vertices - 1; i >= 0; i-- {
			fmt.Fprint(fout, answer[i]+1, " ")
		}
	}
}
