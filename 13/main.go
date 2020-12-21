package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getLines(filename string) ([]string, error) {
	var lines []string

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		raw := scanner.Text()

		lines = append(lines, raw)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func main() {
	lines, err := getLines("in")
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	earliest, err := strconv.Atoi(lines[0])
	if err != nil {
		fmt.Printf("error %v parsing number: %s\n", err, lines[0])
	}

	bestBus := -1
	bestWait := -1
	buses := strings.Split(lines[1], ",")
	for _, bus := range buses {
		if bus == "x" {
			continue
		}
		num, err := strconv.Atoi(bus)
		if err != nil {
			fmt.Printf("error %v parsing number: %s\n", err, bus)
		}
		wait := num - (earliest % num)
		if bestWait == -1 || wait < bestWait {
			bestWait = wait
			bestBus = num
		}
	}

	fmt.Printf("%d %d = %d\n", bestBus, bestWait, bestBus*bestWait)

	// 	// map of time since t to bus to check for
	// 	checks := make(map[int]int)
	// 	for i, bus := range buses {
	// 		if bus == "x" {
	// 			continue
	// 		}
	// 		num, err := strconv.Atoi(bus)
	// 		if err != nil {
	// 			fmt.Printf("error %v parsing number: %s\n", err, bus)
	// 		}
	// 		checks[i] = num
	// 	}
	// 	fmt.Printf("checks %#v\n", checks)

	// outer:
	// 	// 100000000000000
	// 	for t := int64(5589714924) * int64(1789); ; t += int64(checks[0]) {
	// 		for i, bus := range checks {
	// 			current := t + int64(i)
	// 			remainder := current % int64(bus)
	// 			if remainder != 0 {
	// 				// fmt.Printf("%d %d\n", current, bus)
	// 				// t += remainder
	// 				continue outer
	// 			}
	// 		}

	// 		// success!
	// 		fmt.Printf("%d\n", t)
	// 		break
	// 	}
}
