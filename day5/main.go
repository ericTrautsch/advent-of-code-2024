package main

import (
	"fmt"
)

func main() {
	fmt.Println("day5!!")
}

func buildDependencies(rules [][]int) map[int]int {
	return map[int]int{2: 3}
}

func isValidUpdate(update []int, rules map[int]int) bool {
	return false
}

func getMiddleValue(update []int) int {
	return update[len(update)/2]
}

func readRules(path string) [][]int {
	return [][]int{{3}}
}

func readUpdates(path string) [][]int {
	return [][]int{{3}}
}
