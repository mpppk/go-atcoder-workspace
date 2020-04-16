package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64, M int64, X []int64, Y []int64, Z []int64) string {
	return ""
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var N int64
	scanner.Scan()
	N, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var M int64
	scanner.Scan()
	M, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	X := make([]int64, M)
	Y := make([]int64, M)
	Z := make([]int64, M)
	for i := int64(0); i < M; i++ {
		scanner.Scan()
		X[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		Y[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		Z[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	fmt.Println(solve(N, M, X, Y, Z))
}
