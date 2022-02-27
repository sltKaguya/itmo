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
	split := strings.Fields((scanner.Text()))
	size, _ := strconv.Atoi(split[0])

	matrix := make([][]int, size)
	for i := range matrix {
		matrix[i] = make([]int, size)
	}

	for i := 0; i < size; i++ {
		scanner.Scan()
		line := strings.Fields(scanner.Text())
		for y := 0; y < size; y++ {
			matrix[y][i], _ = strconv.Atoi(line[y])
		}
	}

	fout, _ := os.Create("output.txt")
	check := 1

	for i := 0; i < size; i++ {
		for y := i; y < size; y++ {
			if matrix[y][i] != matrix[i][y] {
				fmt.Fprintln(fout, "NO")
				check = 0
				break
			}
		}

		if matrix[i][i] == 1 {
			fmt.Fprintln(fout, "NO")
			check = 0
			break
		}
	}

	if check == 1 {
		fmt.Fprintln(fout, "YES")
	}
}
