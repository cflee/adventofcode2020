package main

import (
	"bufio"
	"fmt"
	"os"
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

// Instruction is one instruction in the code
type instruction struct {
	op  string
	arg int
}

func parseInstructions(lines []string) ([]instruction, error) {
	var output []instruction

	for _, line := range lines {
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			return output, fmt.Errorf("instruction instead of 2 parts had: %d", len(parts))
		}

		var arg int
		_, err := fmt.Sscanf(parts[1], "%d", &arg)
		if err != nil {
			return output, fmt.Errorf("could not parse arg: %s", parts[1])
		}

		output = append(output, instruction{
			op:  parts[0],
			arg: arg,
		})
	}

	return output, nil
}

func execute(inst []instruction) (int, error) {
	seen := make(map[int]bool)
	ip := 0
	acc := 0
	for ip < len(inst) {
		// break and report before executing any instruction a second time
		if seen[ip] {
			return acc, fmt.Errorf("executing ip %d again", ip)
		}
		seen[ip] = true

		cur := inst[ip]
		switch cur.op {
		case "nop":
			ip++
		case "acc":
			acc += cur.arg
			ip++
		case "jmp":
			ip += cur.arg
		}
	}
	return acc, nil
}

func main() {
	lines, err := getLines("in")
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	inst, err := parseInstructions(lines)
	if err != nil {
		fmt.Printf("error parsing instructions: %v\n", err)
	}

	// part one
	acc, err := execute(inst)
	if err == nil {
		fmt.Printf("there should be an error for this!\n")
	}
	fmt.Printf("acc: %d\n", acc)

	// part two: could slightly reduce this by identifying the executed jmp
	// and nop in the earlier run as candidates, but let's just brute force
	for i := 0; i < len(inst); i++ {
		if inst[i].op == "acc" {
			continue
		}

		// flip the instruction
		if inst[i].op == "nop" {
			inst[i].op = "jmp"
		} else {
			inst[i].op = "nop"
		}

		acc, err := execute(inst)
		if err == nil {
			// success!
			fmt.Printf("acc: %d\n", acc)
		}

		// flip the instruction back
		if inst[i].op == "nop" {
			inst[i].op = "jmp"
		} else {
			inst[i].op = "nop"
		}
	}
}
