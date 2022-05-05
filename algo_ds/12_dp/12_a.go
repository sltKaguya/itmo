package main

import (
	"bufio"
	"fmt"
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

	dp := make([]int, n)     //Lenght of the largest increasing subsequence ending by i-th element
	parent := make([]int, n) //Index of previous elem in the LIS

	for i := 0; i < n; i++ {
		dp[i], parent[i] = 1, -1
		for j := 0; j < i; j++ {
			if sequence[j] < sequence[i] {
				if dp[i] < dp[j]+1 {
					dp[i] = dp[j] + 1
					parent[i] = j
				}
			}
		}
	}

	k := 0
	ind := -1
	for i := range dp {
		if dp[i] > k {
			k = dp[i]
			ind = i
		}
	}

	answer := make([]int, k)

	for i := range answer {
		answer[i] = sequence[ind]
		ind = parent[ind]
	}

	fmt.Println(k)
	for i := len(answer) - 1; i >= 0; i-- {
		fmt.Print(answer[i], " ")
	}
}
