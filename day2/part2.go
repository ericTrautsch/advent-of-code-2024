package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func read_reports(path string) [][]int {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error: couldn't read reports", err)
	}

	defer file.Close()

	reports := [][]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		report := []int{}
		text := strings.Split(scanner.Text(), " ")

		for _, str := range text {
			num, err := strconv.Atoi(str)
			if err != nil {
				fmt.Println("Error converting string to int", err)
			}
			report = append(report, num)
		}
		reports = append(reports, report)
	}
	return reports
}

func is_valid_report(report []int) bool {
	// Must be all increasing
	all_increasing := true
	// Or all decreasing
	all_decreasing := true
	// Levels must differ between 1-3
	for i := range report {
		// if it makes it to the end, skip out the loop
		if i+1 == len(report) {
			return all_decreasing || all_increasing
		}
		difference := report[i] - report[i+1]
		// Negative escape condition, handles the equal case
		if int(math.Abs(float64(difference))) > 3 || int(math.Abs(float64(difference))) < 1 {
			return false
		}

		if difference < 0 {
			all_decreasing = false
		} else {
			all_increasing = false
		}
	}
	return all_decreasing || all_increasing
}

func dampener(report []int) bool {
	if is_valid_report(report) {
		return true
	}
	for i := 0; i < len(report); i++ {
		// Create a new array excluding the element at index i
		heldOut := report[i]
		// rest := append(report[:i], report[i+1:]...)
		// Get the rest of the array
		rest := make([]int, 0, len(report)-1) // Preallocate space for performance
		rest = append(rest, report[:i]...)    // Append elements before the current index
		rest = append(rest, report[i+1:]...)  // Append elements after the current index

		fmt.Printf("Held out: %d, Rest: %v\n", heldOut, rest)
		if is_valid_report(rest) {
			return true
		}
	}
	//
	//	for i := 0; i < len(report); i++ {
	//		// fmt.Println(append(report[:i], report[(i+1):]...))
	//		start := report[:i]
	//		end := report[i+1:]
	//		fmt.Println(start, "start")
	//		fmt.Println(end, "end")
	//		if is_valid_report(append(start, end...)) {
	//			fmt.Println("Valid!")
	//		}
	//	}
	return false
}

func main() {
	filepath := "input.txt"
	fmt.Println(len(read_reports(filepath)))
	fmt.Println((read_reports(filepath)))
	count := 0

	for _, report := range read_reports(filepath) {
		if dampener(report) {
			count++
		}
	}
	fmt.Println(count)
}
