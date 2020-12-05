package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getInput(filename string) ([]string, error) {
	var passes []string

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		raw := scanner.Text()

		passes = append(passes, raw)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return passes, nil
}

func parsePass(pass string) (int, int) {
	p := []byte(pass)
	for i := 0; i < len(p); i++ {
		switch p[i] {
		case 'F':
			p[i] = '0'
		case 'B':
			p[i] = '1'
		case 'L':
			p[i] = '0'
		case 'R':
			p[i] = '1'
		}
	}
	s := string(p)

	r, err := strconv.ParseInt(s[:7], 2, 8)
	if err != nil {
		return 0, 0
	}
	c, err := strconv.ParseInt(s[7:], 2, 8)
	if err != nil {
		return 0, 0
	}

	// row 0-127 can fit in an int8?
	return int(r), int(c)
}

func main() {
	passes, err := getInput("in")
	if err != nil {
		fmt.Printf("%v", err)
	}

	max := 0
	occupied := make(map[int]bool, 128*8)
	for _, pass := range passes {
		r, c := parsePass(pass)
		seatid := (r * 8) + c

		if seatid > max {
			max = seatid
		}
		occupied[seatid] = true
	}

	fmt.Printf("%d\n", max)

	// should do some binary search here
	// either doing passes to locate the occupied region, then third pass within
	// or maybe a single-pass modified to consider the occupied region
	// but it's only 128*8 seats...
	pos := 0
	// go to the first occupied seat
	for pos < 128*8 {
		if occupied[pos] == true {
			break
		}
		pos++
	}
	// go to the first unoccupied seat from the first occupied seat
	for pos < 128*8 {
		if occupied[pos] == false {
			break
		}
		pos++
	}
	// done!
	fmt.Printf("%d\n", pos)
}
