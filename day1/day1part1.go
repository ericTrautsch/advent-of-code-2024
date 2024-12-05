package p2

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func _read_lists(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return data
	}
	return data
}

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

	total_distance := 0
	for i := 0; i < len(list_one); i++ {
		total_distance += int(math.Abs(float64(list_one[i] - list_two[i])))
	}

	fmt.Println(total_distance)
}
