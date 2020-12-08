package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type link struct {
	qty  int
	name string
}

type bag struct {
	name        string
	containedBy []string
	contains    []link
}

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

func parseLines(lines []string) (map[string]bag, error) {
	bags := make(map[string]bag)

	for _, line := range lines {
		words := strings.Split(line, " ")
		if len(words) < 7 {
			// min length is "x y bags contain no other bags." (7)
			return bags, fmt.Errorf("line is too short: %s", line)
		}

		// get or init new bag node
		name := fmt.Sprintf("%s %s", words[0], words[1])
		bag, ok := bags[name]
		if !ok {
			bag.name = name
		}

		for i := 4; i+3 < len(words); i += 4 {
			// "no other bags." => i on "no" => i+3 exceeds length
			// so there's no need to specially catch it

			// parse out the qty
			qty, err := strconv.Atoi(words[i])
			if err != nil {
				return bags, fmt.Errorf("unable to parse qty: %s", words[i])
			}

			// establish the down link
			downBagName := fmt.Sprintf("%s %s", words[i+1], words[i+2])
			bag.contains = append(bag.contains, link{
				qty:  qty,
				name: downBagName,
			})

			// up link back from down linked
			downBag := bags[downBagName]
			// assume their init name will be done at some other point...
			downBag.containedBy = append(downBag.containedBy, name)
			bags[downBagName] = downBag
		}

		bags[name] = bag
	}

	return bags, nil
}

func main() {
	lines, err := getLines("in")
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	bags, err := parseLines(lines)

	// traverse up to find distinct bags that could contain shiny gold bag
	seen := make(map[string]bool)
	queue := bags["shiny gold"].containedBy
	for len(queue) > 0 {
		elem := queue[0]
		queue = queue[1:]

		if seen[elem] == true {
			continue
		}
		seen[elem] = true
		for _, c := range bags[elem].containedBy {
			queue = append(queue, c)
		}
	}
	fmt.Printf("%d\n", len(seen))

	// traverse down to find total number of bags inside a shiny gold bag
	sum := 0
	downQueue := bags["shiny gold"].contains
	for len(downQueue) > 0 {
		l := downQueue[0]
		downQueue = downQueue[1:]

		if len(bags[l.name].contains) == 0 {
			// end of the chain!
			sum += l.qty
			continue
		}

		// count this current bag
		sum += l.qty

		for _, l2 := range bags[l.name].contains {
			downQueue = append(downQueue, link{
				qty:  l.qty * l2.qty,
				name: l2.name,
			})
		}
	}
	fmt.Printf("%d\n", sum)
}
