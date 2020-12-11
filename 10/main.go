package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func getNums(filename string) ([]int, error) {
	var lines []int

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		raw := scanner.Text()
		u, err := strconv.ParseInt(raw, 10, 0)
		if err != nil {
			return nil, err
		}

		lines = append(lines, int(u))
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func main() {
	nums, err := getNums("in")
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	sort.Ints(nums)

	// zero to the first adapter
	diff1 := nums[0]
	// last adapter to the device
	diff3 := 1

	for i := 1; i < len(nums); i++ {
		switch nums[i] - nums[i-1] {
		case 1:
			diff1++
		case 3:
			diff3++
		}
	}

	fmt.Printf("%d x %d = %d\n", diff1, diff3, diff1*diff3)

	// need to add the initial 0 to handle 1,2,3 adapters
	nums = append([]int{0}, nums...)
	// tail device doesn't matter since +3 means no extra paths

	t := make([]int, len(nums))
	t[0] = 1 // seed the initial
	for i := 1; i < len(nums); i++ {
		sum := 0
		if i-3 >= 0 && nums[i-3]+3 >= nums[i] {
			sum = sum + t[i-3]
		}
		if i-2 >= 0 && nums[i-2]+3 >= nums[i] {
			sum = sum + t[i-2]
		}
		if i-1 >= 0 && nums[i-1]+3 >= nums[i] {
			sum = sum + t[i-1]
		}
		t[i] = sum
	}
	fmt.Printf("%d\n", t[len(t)-1])
}
