package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fin, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	split := strings.Fields(scanner.Text())
	vertices, _ := strconv.Atoi(split[0])

	matrix := make([][]int, vertices)
	for i := range matrix {
		matrix[i] = make([]int, vertices)
	}

	fout, _ := os.Create("output.txt")
	check := 1

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		x, _ := strconv.Atoi(line[0])
		y, _ := strconv.Atoi(line[1])
		matrix[x-1][y-1] += 1
		if matrix[x-1][y-1] > 1 {
			fmt.Fprintf(fout, "YES")
			check = 0
			break
		}
		matrix[y-1][x-1] += 1
	}

	if check == 1 {
		fmt.Fprintf(fout, "NO")
	}
}
