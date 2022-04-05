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
	degree string
	first  *Node
}

type Graph struct {
	adjList []LinkedList
}

func (l *LinkedList) pushNode(vertex int) {
	l.first = &Node{vertex: vertex, next: l.first}
}

func (l *LinkedList) countDeg() {
	deg := 0
	elem := l.first

	for elem != nil {
		deg++
		elem = elem.next
	}

	l.degree = strconv.Itoa(deg)
}

func main() {
	fin, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	buffer := strings.Fields(scanner.Text())
	vertices, _ := strconv.Atoi(buffer[0])

	myGraph := &Graph{adjList: make([]LinkedList, vertices)}

	for scanner.Scan() {
		buffer = strings.Fields(scanner.Text())
		x, _ := strconv.Atoi(buffer[0])
		y, _ := strconv.Atoi(buffer[1])
		myGraph.adjList[x-1].pushNode(y - 1)
		myGraph.adjList[y-1].pushNode(x - 1)
	}

	for i := range myGraph.adjList {
		myGraph.adjList[i].countDeg()
	}

	fout, _ := os.Create("output.txt")
	for i := range myGraph.adjList {
		fmt.Fprint(fout, myGraph.adjList[i].degree, " ")
	}
}
