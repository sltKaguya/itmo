package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	vertex   int
	next     *Node
	previous *Node
}

type LinkedList struct {
	color    string
	parent   int
	comp_num int
	first    *Node
}

func (list *LinkedList) putNode(vert int) {
	list.color = "white"
	elem := &Node{
		vertex: vert,
	}

	if list.first != nil {
		list.first.previous = elem
	}
	elem.next = list.first
	list.first = elem
}

func (list *LinkedList) DfsVisit(adjList []LinkedList, comp_num int) {
	list.color = "gray"
	list.comp_num = comp_num
	elem := list.first

	for elem != nil {
		if adjList[elem.vertex].color == "white" {
			adjList[elem.vertex].parent = elem.vertex
			adjList[elem.vertex].comp_num = comp_num
			list.DfsVisit(adjList, comp_num)
		}
	}

	list.color = "black"
}

func main() {
	fin, _ := os.Open("components.in")
	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	split := strings.Fields(scanner.Text())
	vertices, _ := strconv.Atoi(split[0])
	adjList := make([]LinkedList, vertices)
	comp_num := 0

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		x, _ := strconv.Atoi(line[0])
		x--
		y, _ := strconv.Atoi(line[1])
		y--
		adjList[x].putNode(y)
		adjList[y].putNode(x)
	}

	for i := range adjList {
		if adjList[i].color == "white" {
			comp_num++
			adjList[i].DfsVisit(adjList, comp_num)
		}
	}

	fmt.Println(adjList)
	fout, _ := os.Create("components.out")
	fmt.Fprintln(fout, comp_num)
	for _, elem := range adjList {
		fmt.Fprint(fout, elem.comp_num)
		fmt.Fprint(fout, " ")
	}
}
