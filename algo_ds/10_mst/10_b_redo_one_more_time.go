package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fin, _ := os.Open("spantree.in")
	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	buffer := strings.Fields(scanner.Text())
	N, _ := strconv.Atoi(buffer[0])
	vertices := make([][]int, N)

	for i := 0; i < N; i++ {
		scanner.Scan()
		buffer = strings.Fields(scanner.Text())
		vertices[i] = make([]int, 2)
		vertices[i][0], _ = strconv.Atoi(buffer[0])
		vertices[i][1], _ = strconv.Atoi(buffer[1])
	}

	myGraph := make([][]float64, N)
	for i := range myGraph {
		myGraph[i] = make([]float64, N)
	}

	for i := range myGraph {
		for j := range myGraph[i] {
			myGraph[i][j] = math.Sqrt(math.Pow(float64(vertices[j][0]-vertices[i][0]), 2) + math.Pow(float64(vertices[j][1]-vertices[i][1]), 2))
			myGraph[j][i] = myGraph[i][j]
		}
	}

	nodes := make([]float64, N)
	for i := 0; i < N; i++ {
		nodes[i] = math.Inf(1)
	}
	visited := make([]float64, N)
	for i := 0; i < N; i++ {
		visited[i] = 0
	}

	var MSTWeight float64 = 0
	nodes[0] = 0

	for i := 0; i < N; i++ {
		vertex := -1
		for j := 0; j < N; j++ {
			if visited[j] == 0 && (vertex == -1 || nodes[j] < nodes[vertex]) {
				vertex = j
			}
		}
		visited[vertex] = 1
		for j := 0; j < N; j++ {
			if visited[j] == 0 && myGraph[vertex][j] < nodes[j] && vertex != j {
				nodes[j] = myGraph[vertex][j]
			}
		}
	}

	for i := 0; i < N; i++ {
		MSTWeight += nodes[i]
	}
	fout, _ := os.Create("spantree.out")
	fmt.Fprint(fout, MSTWeight)
}
