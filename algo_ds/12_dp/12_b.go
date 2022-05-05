package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	buffer := strings.Fields(scanner.Text())
	n, _ := strconv.Atoi(buffer[0])

	scanner.Scan()
	buffer = strings.Fields(scanner.Text())
	sequence := make([]int, n)
	for i := 0; i < n; i++ {
		sequence[i], _ = strconv.Atoi(buffer[i])
	}

	dp := make([]int, n)
	pos := make([]int, n)
	prev := make([]int, n)

	k := 0
	pos[0] = -1
}
