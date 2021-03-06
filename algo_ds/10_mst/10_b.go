package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	vertex int
	weight float64
	next   *Node
}

type LinkedList struct {
	first *Node
	dist  float64
	done  bool //if vertice in Queue - not done, once extracted - done
}

type Graph struct {
	adjList []LinkedList
}

func (list *LinkedList) pushNode(vertex int, weight float64) {
	list.first = &Node{vertex: vertex, weight: weight, next: list.first}
}

func (g *Graph) findMin() int {
	minIndex := -1
	min := math.Inf(1)
	for i, elem := range g.adjList {
		if elem.dist < min {
			minIndex = i
			min = elem.dist
		}
	}

	return minIndex
}

func (l *LinkedList) findMin(g *Graph) float64 {
	min := math.Inf(1)
	elem := l.first
	for elem != nil {
		if elem.weight < min && !g.adjList[elem.vertex].done {
			min = elem.weight
		}

		elem = elem.next
	}

	return min
}

func main() {
	fin, _ := os.Open("spantree.in")
	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	buffer := strings.Fields(scanner.Text())
	nVertices, _ := strconv.Atoi(buffer[0])

	vertices := make([][]float64, nVertices)
	for i := range vertices {
		vertices[i] = make([]float64, 2)
	}

	i := 0
	for scanner.Scan() {
		buffer = strings.Fields(scanner.Text())
		vertices[i][0], _ = strconv.ParseFloat(buffer[0], 64)
		vertices[i][1], _ = strconv.ParseFloat(buffer[1], 64)
		i++
	}

	myGraph := &Graph{adjList: make([]LinkedList, nVertices)}
	for i := range myGraph.adjList {
		myGraph.adjList[i].dist = math.Inf(1)
	}

	for i, elemOne := range vertices {
		for j, elemTwo := range vertices {
			if i != j {
				weight := math.Sqrt(math.Pow(elemTwo[0]-elemOne[0], 2) + math.Pow(elemTwo[1]-elemOne[1], 2))
				myGraph.adjList[i].pushNode(j, weight)
			}
		}
	}

	myGraph.adjList[0].dist = 0
	myGraph.adjList[0].done = true

	MSTCount := 0
	MSTWeight := 0

	for MSTCount < nVertices {
		i := myGraph.findMin()
		elem := myGraph.adjList[i].first
		for elem != nil {
			myGraph.adjList[elem.vertex].dist = elem.weight
		}
		MSTWeight += int(myGraph.adjList[i].findMin(myGraph))
		myGraph.adjList[i].done = true
		MSTCount++
	}

	for i, elem := range myGraph.adjList {
		fmt.Println("BOBA", i, elem.dist)
		selem := elem.first
		for selem != nil {
			fmt.Println("biba", selem.vertex, " ", selem.weight)
			selem = selem.next
		}
	}

	fmt.Println(MSTWeight)

	//fout, _ := os.Create("spantree.out")
}
