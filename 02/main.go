package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Line struct {
	lower    int
	upper    int
	letter   byte
	password string
}

func getInput(filename string) ([]Line, error) {
	var lines []Line

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		raw := scanner.Text()
		s1 := strings.Split(raw, " ")

		qtys := strings.Split(s1[0], "-")
		lower, err := strconv.Atoi(qtys[0])
		if err != nil {
			return nil, fmt.Errorf("unable to parse qty lower %w", err)
		}
		upper, err := strconv.Atoi(qtys[1])
		if err != nil {
			return nil, fmt.Errorf("unable to parse qty upper %w", err)
		}

		letter := s1[1][0]

		password := s1[2]

		line := Line{
			lower:    lower,
			upper:    upper,
			letter:   letter,
			password: password,
		}

		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func main() {
	lines, err := getInput("in")
	if err != nil {
		fmt.Printf("%v", err)
	}

	valid := 0
	for _, line := range lines {
		// let's just treat it as bytes and not runes...
		count := 0
		for i := 0; i < len(line.password); i++ {
			if line.password[i] == line.letter {
				count++
			}
		}

		if count >= line.lower && count <= line.upper {
			valid++
		}
	}

	fmt.Printf("%v\n", valid)

	valid = 0
	for _, line := range lines {
		// let's just treat it as bytes and not runes...
		pos1 := line.password[line.lower-1] == line.letter
		pos2 := line.password[line.upper-1] == line.letter

		if pos1 != pos2 {
			valid++
		}
	}

	fmt.Printf("%v\n", valid)
}
