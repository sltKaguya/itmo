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

type Graph struct {
	adjList []*Node
	dist    []int
	state   []bool
}

type Queue struct {
	first *Node
	last  *Node
}

func (queue *Queue) push(vert int) {
	elem := &Node{vertex: vert}
	if queue.first != nil {
		queue.last.next = elem
		queue.last = elem
	} else {
		queue.last = elem
		queue.first = queue.last
	}
}

func (queue *Queue) pop() *Node {
	elem := queue.first
	queue.first = queue.first.next
	return elem
}

func (queue *Queue) isEmpty() bool {
	return queue.first == nil
}

func main() {
	fin, _ := os.Open("pathbge1.in")
	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	line := strings.Fields(scanner.Text())
	vertices, _ := strconv.Atoi(line[0])

	myGraph := &Graph{
		state:   make([]bool, vertices),
		adjList: make([]*Node, vertices),
		dist:    make([]int, vertices),
	}

	for scanner.Scan() {
		line = strings.Fields(scanner.Text())
		x, _ := strconv.Atoi(line[0])
		y, _ := strconv.Atoi(line[1])
		myGraph.adjList[x-1] = &Node{vertex: y - 1, next: myGraph.adjList[x-1]}
		myGraph.adjList[y-1] = &Node{vertex: x - 1, next: myGraph.adjList[y-1]}
	}

	myQueue := Queue{}
	myQueue.push(0)
	myGraph.state[0] = true

	for !myQueue.isEmpty() {
		elem := myQueue.pop()
		node := myGraph.adjList[elem.vertex]
		for node != nil {
			if !myGraph.state[node.vertex] {
				myQueue.push(node.vertex)
				myGraph.state[node.vertex] = true
				myGraph.dist[node.vertex] = myGraph.dist[elem.vertex] + 1
			}
			node = node.next
		}
	}

	fout, _ := os.Create("pathbge1.out")

	for _, elem := range myGraph.dist {
		fmt.Fprint(fout, elem, " ")
	}
}
