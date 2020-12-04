package main

import (
	"bufio"
	"fmt"
	"os"
)

func getInput(filename string) ([][]byte, error) {
	var grid [][]byte

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		raw := scanner.Text()
		row := []byte(raw)

		grid = append(grid, row)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return grid, nil
}

func slope(grid [][]byte, right int, down int) int {
	trees := 0
	for row := 0; (row * down) < len(grid); row++ {
		col := (row * right) % len(grid[row*down])
		if grid[row*down][col] == '#' {

			trees++
		}
	}
	return trees
}

func main() {
	grid, err := getInput("in")
	if err != nil {
		fmt.Printf("%v", err)
	}

	trees := slope(grid, 3, 1)
	fmt.Printf("%d\n", trees)

	trees = slope(grid, 1, 1) * slope(grid, 3, 1) * slope(grid, 5, 1) * slope(grid, 7, 1) * slope(grid, 1, 2)
	fmt.Printf("%d\n", trees)
}
