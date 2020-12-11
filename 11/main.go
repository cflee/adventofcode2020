package main

import (
	"bufio"
	"fmt"
	"os"
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

var neighbours = [][]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

func main1() {
	lines, err := getLines("in")
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	map1 := make([][]byte, len(lines))
	map2 := make([][]byte, len(lines))
	for i, line := range lines {
		map1[i] = []byte(line)
		map2[i] = []byte(line)
	}

	for {
		changed := false

		for i := 0; i < len(map1); i++ {
			for j := 0; j < len(map1[i]); j++ {
				// floor never changes
				if map1[i][j] == '.' {
					continue
				}

				occupied := 0
				for _, n := range neighbours {
					x := i + n[0]
					y := j + n[1]
					if x >= 0 && x < len(map1) && y >= 0 && y < len(map1[i]) {
						if map1[x][y] == '#' {
							occupied++
						}
					}
				}

				switch {
				case map1[i][j] == 'L' && occupied == 0:
					map2[i][j] = '#'
					changed = true
				case map1[i][j] == '#' && occupied >= 4:
					map2[i][j] = 'L'
					changed = true
				default:
					map2[i][j] = map1[i][j]
				}
			}
		}

		// swap refs to 1 and 2
		temp := map1
		map1 = map2
		map2 = temp

		if !changed {
			break
		}
	}

	occupied := 0
	for i := 0; i < len(map1); i++ {
		for j := 0; j < len(map1[i]); j++ {
			if map1[i][j] == '#' {
				occupied++
			}
		}
	}
	fmt.Printf("%d\n", occupied)
}

func main2() {
	lines, err := getLines("in")
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	map1 := make([][]byte, len(lines))
	map2 := make([][]byte, len(lines))
	for i, line := range lines {
		map1[i] = []byte(line)
		map2[i] = []byte(line)
	}

	for {
		changed := false

		for i := 0; i < len(map1); i++ {
			for j := 0; j < len(map1[i]); j++ {
				// floor never changes
				if map1[i][j] == '.' {
					continue
				}

				occupied := 0
				for _, n := range neighbours {
					x := i + n[0]
					y := j + n[1]

				l1:
					for x >= 0 && x < len(map1) && y >= 0 && y < len(map1[i]) {
						switch map1[x][y] {
						case '#':
							occupied++
							break l1
						case 'L':
							break l1
						default:
							x = x + n[0]
							y = y + n[1]
						}

					}
				}

				switch {
				case map1[i][j] == 'L' && occupied == 0:
					map2[i][j] = '#'
					changed = true
				case map1[i][j] == '#' && occupied >= 5:
					map2[i][j] = 'L'
					changed = true
				default:
					map2[i][j] = map1[i][j]
				}
			}
		}

		// swap refs to 1 and 2
		temp := map1
		map1 = map2
		map2 = temp

		if !changed {
			break
		}
	}

	occupied := 0
	for i := 0; i < len(map1); i++ {
		for j := 0; j < len(map1[i]); j++ {
			if map1[i][j] == '#' {
				occupied++
			}
		}
	}
	fmt.Printf("%d\n", occupied)
}

func main() {
	main1()
	main2()
}
