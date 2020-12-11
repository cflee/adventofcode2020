package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getNums(filename string) ([]uint64, error) {
	var lines []uint64

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		raw := scanner.Text()
		u, err := strconv.ParseUint(raw, 10, 64)
		if err != nil {
			return nil, err
		}

		lines = append(lines, u)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func main1(size int) uint64 {
	nums, err := getNums("in")
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	window := make(map[uint64]int)
	// init the preamble
	for i := 0; i < size; i++ {
		window[nums[i]] = window[nums[i]] + 1
	}

	for i := size; i < len(nums); i++ {
		// look for match
		found := false
		for x := uint64(1); x < (nums[i]/2 + 1); x++ {
			if window[x] > 0 && window[nums[i]-x] > 0 {
				found = true
			}
		}
		if !found {
			return nums[i]
		}

		// set up for next:
		// remove the first elem in window
		// and add this current item as last elem in new window
		window[nums[i-size]] = window[nums[i-size]] - 1
		window[nums[i]] = window[nums[i]] + 1
	}

	return 0
}

func main2(goal uint64) uint64 {
	nums, err := getNums("in")
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	left := 0
	right := 1
	sum := nums[0] + nums[1]
	for {
		if sum < goal {
			right++
			sum = sum + nums[right]
		} else if sum > goal {
			sum = sum - nums[left]
			left++
		} else {
			break
		}
	}

	min := nums[left]
	max := nums[left]
	for i := left; i <= right; i++ {
		if nums[i] < min {
			min = nums[i]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}

	return min + max
}

func main() {
	val := main1(25)
	fmt.Printf("found: %d\n", val)
	sum := main2(90433990)
	fmt.Printf("sum: %d\n", sum)
}
