package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64, u []int64, v []int64, w []int64) string {
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
	u := make([]int64, N-1)
	v := make([]int64, N-1)
	w := make([]int64, N-1)
	for i := int64(0); i < N-1; i++ {
		scanner.Scan()
		u[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		v[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		w[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	fmt.Println(solve(N, u, v, w))
}
