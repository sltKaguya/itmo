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
	vertex, _ := strconv.Atoi(split[0])

	matrix := make([][]int, vertex)
	for i := range matrix {
		matrix[i] = make([]int, vertex)
	}

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		v_out, _ := strconv.Atoi(line[0])
		v_in, _ := strconv.Atoi(line[1])
		matrix[v_out-1][v_in-1] = 1
	}
	fout, _ := os.Create("output.txt")
	for _, elem := range matrix {
		temp := make([]string, vertex)
		for i, num := range elem {
			temp[i] = fmt.Sprint(num)
		}
		fmt.Fprintln(fout, strings.Join(temp, " "))
	}
}
