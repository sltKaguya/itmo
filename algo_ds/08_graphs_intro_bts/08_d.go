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
	state    bool
	comp_num int
	first    *Node
}

type Graph struct {
	adjList []LinkedList
}

func (list *LinkedList) putNode(vert int) {
	list.first = &Node{vertex: vert, next: list.first}
}

func (g *Graph) DfsVisit(comp_num int, index int) {
	g.adjList[index].state = true
	g.adjList[index].comp_num = comp_num
	elem := g.adjList[index].first

	for elem != nil {
		if !g.adjList[elem.vertex].state {
			//fmt.Println(elem, "elem")
			//fmt.Println(g.adjList[elem.vertex])
			g.DfsVisit(comp_num, elem.vertex)
		}
		elem = elem.next
		// fmt.Println(elem, "elem")
	}
}

func main() {
	fin, _ := os.Open("components.in")
	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	line := strings.Fields(scanner.Text())
	vertices, _ := strconv.Atoi(line[0])

	graph := &Graph{}
	graph.adjList = make([]LinkedList, vertices)
	comp_num := 0

	// fmt.Println(graph.adjList)

	for scanner.Scan() {
		line = strings.Fields(scanner.Text())
		x, _ := strconv.Atoi(line[0])
		y, _ := strconv.Atoi(line[1])
		graph.adjList[x-1].putNode(y - 1)
		graph.adjList[y-1].putNode(x - 1)
	}

	for i := range graph.adjList {
		if !graph.adjList[i].state {
			comp_num++
			graph.DfsVisit(comp_num, i)
		}
	}

	// fmt.Println(graph.adjList)
	fout, _ := os.Create("components.out")
	fmt.Fprintln(fout, comp_num)
	for _, elem := range graph.adjList {
		fmt.Fprint(fout, elem.comp_num)
		fmt.Fprint(fout, " ")
	}
}
