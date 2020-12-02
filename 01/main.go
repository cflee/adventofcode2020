package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getInput(filename string) ([]int, error) {
	var nums []int

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		nums = append(nums, i)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return nums, nil
}

func main() {
	input, err := getInput("in")
	if err != nil {
		fmt.Printf("%v", err)
	}

	target := 2020

	// calculate 2-sum
	seen := make(map[int]bool)
	for _, num := range input {
		check := target - num
		if _, ok := seen[check]; ok {
			fmt.Printf("two-sum: %d\n", (num * check))
			break
		}
		seen[num] = true
	}

	// calculate 3-sum
	// reset the hashtable, seed it with all the input values
	seen = make(map[int]bool)
	for _, num := range input {
		seen[num] = true
	}

	// iterate over the second and third sum, look in the hashtable
outer:
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			check := target - input[i] - input[j]
			if _, ok := seen[check]; ok {
				fmt.Printf("three-sum: %d\n", (input[i] * input[j] * check))
				break outer
			}
		}
	}
}
