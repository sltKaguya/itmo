package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Cell struct {
	state bool
	end   bool
	dist  string
	x     int
	y     int
}

type Node struct {
	cell *Cell
	next *Node
}

type Queue struct {
	first *Node
	last  *Node
}

func (queue *Queue) push(cell *Cell) {
	elem := &Node{cell: cell}
	if queue.first != nil {
		queue.last.next = elem
		queue.last = elem
	} else {
		queue.last = elem
		queue.first = queue.last
	}
}

func (queue *Queue) pop() *Cell {
	elem := queue.first
	queue.first = queue.first.next
	return elem.cell
}

func (queue *Queue) isEmpty() bool {
	return queue.first == nil
}

func main() {
	fin, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	line := strings.Fields(scanner.Text())
	n, _ := strconv.Atoi(line[0])
	m, _ := strconv.Atoi(line[1])

	matrix := make([][]*Cell, n)
	for i := range matrix {
		matrix[i] = make([]*Cell, m)
	}
	var start []int

	for i := 0; i < n; i++ {
		scanner.Scan()
		line := scanner.Text()
		fmt.Println(line)
		for j := 0; j < m; j++ {
			fmt.Println(line[j])
			switch line[j] {
			case '.':
				matrix[i][j] = &Cell{state: true, x: i, y: j}
			case 'T':
				matrix[i][j] = &Cell{end: true, x: i, y: j}
			case 'S':
				matrix[i][j] = &Cell{x: i, y: j}
				start = []int{i, j}
			default:
				matrix[i][j] = &Cell{x: i, y: j}
			}
		}
	}

	myQueue := Queue{}
	myQueue.push(matrix[start[0]][start[1]])
	var endCell *Cell

	for !myQueue.isEmpty() {
		elem := myQueue.pop()
		if elem.end {
			endCell = elem
			break
		}

		if elem.x > 0 {
			leftCell := matrix[elem.x-1][elem.y]
			if leftCell.state && leftCell.dist == "" {
				leftCell.dist = matrix[elem.x][elem.y].dist + "L"
				myQueue.push(leftCell)
			}
		}

		if elem.y > 0 {
			upCell := matrix[elem.x][elem.y-1]
			if upCell.state && upCell.dist == "" {
				upCell.dist = matrix[elem.x][elem.y].dist + "U"
				myQueue.push(upCell)
			}
		}

		if elem.x < m-1 {
			rightCell := matrix[elem.x+1][elem.y]
			if rightCell.state && rightCell.dist == "" {
				rightCell.dist = matrix[elem.x][elem.y].dist + "R"
				myQueue.push(rightCell)
			}
		}

		if elem.x < m-1 {
			downCell := matrix[elem.x][elem.y+1]
			if downCell.state && downCell.dist == "" {
				downCell.dist = matrix[elem.x][elem.y].dist + "D"
				myQueue.push(downCell)
			}
		}
	}

	fout, _ := os.Create("output.txt")
	if endCell == nil {
		fmt.Fprint(fout, "-1")
	} else {
		fmt.Fprintln(fout, len(endCell.dist))
		fmt.Fprint(fout, endCell.dist)
	}
}
