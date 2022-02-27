package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	vertex   int
	next     *Node
	previous *Node
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

type LinkedList struct {
	first    *Node
	parent   int
	color    string
	dist     int
	next     *LinkedList
	prevoius *LinkedList
}

type Queue struct {
	first *LinkedList
}

func (queue *Queue) enQueue(vertex LinkedList) {

}

func main() {
	fin, _ := os.Open("pathbge1.in")
	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	line := strings.Fields(scanner.Text())
	vertices, _ := strconv.Atoi(line[0])

	adjList := make([]LinkedList, vertices)

	for scanner.Scan() {
		line = strings.Fields(scanner.Text())
		x, _ := strconv.Atoi(line[0])
		y, _ := strconv.Atoi(line[1])
		adjList[x-1].putNode(y - 1)
		adjList[y-1].putNode(x - 1)
	}
	for i := range adjList {
		adjList[i].color = "white"
		adjList[i].dist = int(math.Inf(1))
	}
	adjList[0].color = "gray"
	adjList[0].dist = 0

	var myQueue LinkedList

	myQueue.putNode()
}
