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
	done  bool //if vertice not in MST - false, else true
}

type Graph struct {
	adjList []LinkedList
}

func (list *LinkedList) pushNode(vertex int, weight float64) {
	list.first = &Node{vertex: vertex, weight: weight, next: list.first}
}

func (g *Graph) findMin() (float64, int) {
	minIndex := -1
	min := math.Inf(1)
	for i, elem := range g.adjList {
		if elem.dist < min && !elem.done {
			minIndex = i
			min = elem.dist
		}
	}

	return min, minIndex
}

func main() {
	fin, _ := os.Open("spantree.in")
	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	buffer := strings.Fields(scanner.Text())
	countVertices, _ := strconv.Atoi(buffer[0])

	vertices := make([][]float64, countVertices)
	for i := range vertices {
		vertices[i] = make([]float64, 2)
	}

	for i := 0; i < countVertices; i++ {
		scanner.Scan()
		buffer = strings.Fields(scanner.Text())
		vertices[i][0], _ = strconv.ParseFloat(buffer[0], 64)
		vertices[i][1], _ = strconv.ParseFloat(buffer[1], 64)
	}

	myGraph := &Graph{adjList: make([]LinkedList, countVertices)}
	for i := range myGraph.adjList {
		myGraph.adjList[i].dist = math.Inf(1)
	}

	for i, elemOne := range vertices {
		for j, elemTwo := range vertices {
			weight := math.Sqrt(math.Pow(elemTwo[0]-elemOne[0], 2) + math.Pow(elemTwo[1]-elemOne[1], 2))
			myGraph.adjList[i].pushNode(j, weight)
		}
	}

	vertices = nil

	myGraph.adjList[0].dist = 0

	MSTCount := 0
	var MSTWeight float64 = 0

	for MSTCount < countVertices {
		dist, i := myGraph.findMin()
		elem := myGraph.adjList[i].first
		for elem != nil {
			if myGraph.adjList[elem.vertex].dist > elem.weight {
				myGraph.adjList[elem.vertex].dist = elem.weight
			}
			elem = elem.next
		}

		MSTWeight += dist
		myGraph.adjList[i].done = true
		MSTCount++
	}

	fout, _ := os.Create("spantree.out")
	fmt.Fprint(fout, MSTWeight)
}
