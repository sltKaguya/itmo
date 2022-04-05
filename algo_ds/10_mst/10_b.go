package main

import (
	"bufio"
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
	vertices, _ := strconv.Atoi(buffer[0])

	fout, _ := os.Create("spantree.out")
}
