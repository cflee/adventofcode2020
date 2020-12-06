package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	// just read all the bytes instead of line by line
	file, err := ioutil.ReadFile("in")
	if err != nil {
		fmt.Printf("error reading file %v\n", err)
		return
	}

	sum := 0
	groups := strings.Split(string(file), "\n\n")
	for _, group := range groups {
		group = strings.ReplaceAll(group, "\n", "")
		q := make(map[byte]bool)
		for i := 0; i < len(group); i++ {
			q[group[i]] = true
		}
		sum += len(q)
	}
	fmt.Printf("sum of counts %d\n", sum)

	sum = 0
	groups = strings.Split(string(file), "\n\n")
	for _, group := range groups {
		// last group may have phantom person if there's a newline at EOF
		// number of persons is used for checks later
		group = strings.TrimRight(group, "\n")
		persons := strings.Split(group, "\n")

		// assume nobody answers duplicate questions
		seen := make(map[byte]int)
		for _, person := range persons {
			for i := 0; i < len(person); i++ {
				seen[person[i]] = seen[person[i]] + 1
			}
		}
		for _, count := range seen {
			if count == len(persons) {
				sum++
			}
		}
	}
	fmt.Printf("sum of counts %d\n", sum)
}
