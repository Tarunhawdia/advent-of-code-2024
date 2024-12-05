package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// Open the input file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var a, b []int64

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		if len(parts) == 2 {
			x, _ := strconv.ParseInt(parts[0], 10, 64)
			y, _ := strconv.ParseInt(parts[1], 10, 64)
			a = append(a, x)
			b = append(b, y)
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	var ans int64

	countMap := make(map[int64]int64)

	for i := 0; i < len(b); i++ {
		countMap[b[i]]++
	}

	for i := 0; i < len(a); i++ {
		if count, found := countMap[a[i]]; found {
			ans += count * a[i]
		}
	}

	fmt.Println(ans)
}
