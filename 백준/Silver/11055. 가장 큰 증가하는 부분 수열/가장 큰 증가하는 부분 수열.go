package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	N, _ := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)

	arTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var A = make([]int, N)
	for i := 0; i < int(N); i++ {
		arItemTemp, _ := strconv.ParseInt(arTemp[i], 10, 64)
		arItem := int(arItemTemp)
		A[i] = arItem
	}
	var dp = make([]int, N)
	for i := 0; i < int(N); i++ {
		dp = append(dp, 0)
	}
	dp[0] = A[0]

	for i := 0; i < int(N); i++ {
		if i < 1 {
			continue
		}
		for j := 0; j < i; j++ {
			if A[j] < A[i] {
				dp[i] = max(dp[i], dp[j]+A[i])
			} else {
				dp[i] = max(dp[i], A[i])
			}

		}
	}

	maxValue := 0
	for _, v := range dp {
		if v > maxValue {
			maxValue = v
		}
	}
	fmt.Println(maxValue)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
