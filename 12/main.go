package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	"strconv"
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

type position struct {
	x float64 // east
	y float64 // north
}

// move in a polar way
// direction in radians
// math direction starts at 3 o'clock and increases ccw
// navigation bearing/heading starts as 12 o'clock and increases cw
// this is using the math direction sense
func (p *position) move(direction float64, distance float64) {
	p.x = p.x + distance*math.Cos(direction)
	p.y = p.y + distance*math.Sin(direction)
}

// rotate this point around the origin
// positive for cw
func (p *position) rotate(deg float64) {
	// convert to polar way
	distance := math.Sqrt(p.x*p.x + p.y*p.y)
	direction := math.Atan2(p.y, p.x)

	// adjust direction
	direction = direction + (-deg / 180 * math.Pi)
	if direction < -math.Pi {
		direction += 2 * math.Pi
	}
	if direction > math.Pi {
		direction -= 2 * math.Pi
	}

	// then move from origin, towards adjusted direction
	p.x = 0
	p.y = 0
	p.move(direction, distance)
}

func manhattan(pos position) float64 {
	m := float64(0)
	if pos.x < 0 {
		m += -pos.x
	} else {
		m += pos.x
	}
	if pos.y < 0 {
		m += -pos.y
	} else {
		m += pos.y
	}
	return m
}

func q1(lines []string) (position, int) {
	pos := position{}
	dir := int(0)
	for _, line := range lines {
		action := line[0:1]
		param, err := strconv.ParseFloat(string(line[1:]), 64)
		if err != nil {
			fmt.Printf("error parsing number for: %s\n", line)
		}

		switch action {
		case "F":
			pos.move(float64(dir)/180*math.Pi, param)
		case "L":
			dir = (dir + int(param)) % 360
		case "R":
			dir = (dir - int(param)) % 360
		case "N":
			pos.move(float64(90)/180*math.Pi, param)
		case "W":
			pos.move(float64(180)/180*math.Pi, param)
		case "S":
			pos.move(float64(270)/180*math.Pi, param)
		case "E":
			pos.move(float64(0)/180*math.Pi, param)
		}

		// fmt.Printf("%#v dir %d\n", pos, dir)
	}
	return pos, dir
}

func q2(lines []string) (position, int) {
	// waypoint is relative to the ship
	waypoint := position{
		x: 10,
		y: 1,
	}

	// ship is only used as a coordinate
	ship := position{}
	dir := int(0)
	for _, line := range lines {
		action := line[0:1]
		param, err := strconv.ParseFloat(string(line[1:]), 64)
		if err != nil {
			fmt.Printf("error parsing number for: %s\n", line)
		}

		switch action {
		case "F":
			// waypoint.move(float64(dir), param)
			ship.x = ship.x + (param * waypoint.x)
			ship.y = ship.y + (param * waypoint.y)
		case "L":
			waypoint.rotate(-param)
		case "R":
			waypoint.rotate(param)
		case "N":
			waypoint.move(float64(90)/180*math.Pi, param)
		case "W":
			waypoint.move(float64(180)/180*math.Pi, param)
		case "S":
			waypoint.move(float64(270)/180*math.Pi, param)
		case "E":
			waypoint.move(float64(0)/180*math.Pi, param)
		}

		// fmt.Printf("waypoint %#v ship %#v dir %d\n", waypoint, ship, dir)
	}
	return ship, dir
}

func main() {
	lines, err := getLines("in")
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	pos, _ := q1(lines)
	fmt.Printf("pos %#v manhattan distance: %f\n", pos, manhattan(pos))

	ship, _ := q2(lines)
	fmt.Printf("ship %#v manhattan distance %f\n", ship, manhattan(ship))
}
