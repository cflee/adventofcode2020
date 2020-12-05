package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	// not too sure why i can't do a one-pass assign to chars
	// in the string slice so let's just do this
	pass = strings.ReplaceAll(pass, "F", "0")
	pass = strings.ReplaceAll(pass, "B", "1")
	pass = strings.ReplaceAll(pass, "L", "0")
	pass = strings.ReplaceAll(pass, "R", "1")

	r, err := strconv.ParseInt(pass[:7], 2, 8)
	if err != nil {
		return 0, 0
	}
	c, err := strconv.ParseInt(pass[7:], 2, 8)
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
