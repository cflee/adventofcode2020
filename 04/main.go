package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(filename string) ([]string, error) {
	var results []string

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	splitEmptyLine := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		// seek until i, and the previous char, are both newlines
		for i := 1; i < len(data); i++ {
			if data[i-1] == '\n' && data[i] == '\n' {
				// scrub through
				for j := 0; j <= i; j++ {
					if data[j] == '\n' {
						data[j] = ' '
					}
				}
				return i + 1, data[:i], nil
			}
		}

		// otherwise ask for more data
		if !atEOF {
			return 0, nil, nil
		}

		// no more data
		// scrub through
		for j := 0; j < len(data); j++ {
			if data[j] == '\n' {
				data[j] = ' '
			}
		}
		return 0, data, bufio.ErrFinalToken
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(splitEmptyLine)
	for scanner.Scan() {
		raw := scanner.Text()
		// get rid of trailing whitespace
		raw = strings.TrimSpace(raw)
		results = append(results, raw)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

// [min, max] both inclusive
func inRange(str string, min int, max int) bool {
	num, err := strconv.Atoi(str)
	if err != nil {
		return false
	}
	if num < min || num > max {
		return false
	}
	return true
}

func validate(input string) bool {
	data := make(map[string]string)
	fields := strings.Split(input, " ")
	for _, field := range fields {
		kv := strings.Split(field, ":")
		data[kv[0]] = kv[1]
	}

	byr, ok := data["byr"]
	if !ok || len(byr) != 4 || !inRange(byr, 1920, 2002) {
		fmt.Printf("byr %s\n", byr)
		return false
	}

	iyr, ok := data["iyr"]
	if !ok || len(iyr) != 4 || !inRange(iyr, 2010, 2020) {
		fmt.Printf("iyr %s\n", iyr)
		return false
	}

	eyr, ok := data["eyr"]
	if !ok || len(eyr) != 4 || !inRange(eyr, 2020, 2030) {
		fmt.Printf("eyr %s\n", eyr)
		return false
	}

	hgt, ok := data["hgt"]
	if !ok {
		return false
	}
	if strings.HasSuffix(hgt, "cm") {
		if !inRange(hgt[:len(hgt)-2], 150, 193) {
			fmt.Printf("hgt %s\n", hgt)
			return false
		}
	} else if strings.HasSuffix(hgt, "in") {
		if !inRange(hgt[:len(hgt)-2], 59, 76) {
			fmt.Printf("hgt %s\n", hgt)
			return false
		}
	} else {
		fmt.Printf("hgt %s\n", hgt)
		return false
	}

	hcl, ok := data["hcl"]
	if !ok {
		return false
	}
	if len(hcl) != 7 || hcl[0] != '#' {
		fmt.Printf("hcl %s\n", hcl)
		return false
	}
	_, err := strconv.ParseInt(hcl[1:], 16, 25)
	if err != nil {
		fmt.Printf("hcl %s\n", hcl)
		return false
	}

	ecl, ok := data["ecl"]
	if !ok {
		return false
	}
	if ecl != "amb" && ecl != "blu" && ecl != "brn" && ecl != "gry" && ecl != "grn" && ecl != "hzl" && ecl != "oth" {
		fmt.Printf("ecl %s\n", ecl)
		return false
	}

	pid, ok := data["pid"]
	if !ok || len(pid) != 9 || !inRange(pid, 0, 999999999) {
		fmt.Printf("pid %s\n", pid)
		return false
	}

	return true
}

func main() {
	list, err := getInput("in")
	if err != nil {
		fmt.Printf("%v", err)
	}

	valid := 0
	for _, p := range list {
		if validate(p) {
			valid++
		}
	}

	fmt.Printf("%d\n", valid)
}
