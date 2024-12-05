package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func read_lists(path string) []int {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}

	defer file.Close()

	list := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Error converting string:", err)
		}

		list = append(list, num)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
	}

	return list
}

func main() {
	list_one := read_lists("list_one.txt")
	list_two := read_lists("list_two.txt")
	sort.Ints(list_one)
	sort.Ints(list_two)
	freqMap := make(map[int]int)

	for _, item := range list_one {
		freqMap[item] = 0
	}

	for _, item := range list_two {
		freqMap[item]++
	}

	fmt.Println(freqMap)

	total_score := 0
	for i := 0; i < len(list_one); i++ {
		total_score += list_one[i] * freqMap[list_one[i]]
	}

	fmt.Println(total_score)
}
