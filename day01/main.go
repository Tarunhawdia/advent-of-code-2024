package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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

	// Read input from file
	a, b := readInput(file)

	// Solve Part 1
	part1Result := solvePart1(a, b)
	fmt.Println("Part 1 Result:", part1Result)

	// Solve Part 2
	part2Result := solvePart2(a, b)
	fmt.Println("Part 2 Result:", part2Result)
}

// readInput reads and parses the input file into two slices
func readInput(file *os.File) ([]int64, []int64) {
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
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}
	return a, b
}

// solvePart1 calculates the result for Part 1
func solvePart1(a, b []int64) int64 {
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
	var ans int64
	for i := 0; i < len(a); i++ {
		ans += int64(math.Abs(float64(a[i] - b[i])))
	}
	return ans
}

func solvePart2(a, b []int64) int64 {
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

	return ans
}
