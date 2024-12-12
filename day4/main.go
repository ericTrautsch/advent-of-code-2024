package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	word_to_find := "XMAS"
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Errorf("Couldn't read file")
	}
	input := string(data)
	fmt.Println(count_mas(input, word_to_find))
}

func count_mas(corpus string, word string) int {
	lines := strings.Split(corpus, "\n")
	var result [][]string
	for _, line := range lines {
		words := strings.Split(line, "")
		result = append(result, words)
	}
	result = result[:len(result)-1]
	forward, backward := rotate(result)
	// fmt.Println(combine_mas_string(forward))
	// fmt.Println(combine_mas_string(backward))

	return find_mas(combine_mas_string(forward)) + find_mas(combine_mas_string(backward))
}

func find_mas(corpus string) int {
	lines := strings.Split(corpus, "\n")
	var result [][]string
	for _, line := range lines {
		words := strings.Split(line, "")
		result = append(result, words)
	}
	// fmt.Println(result)
	// fmt.Println(len(result))
	count := 0
	for i := 1; i < len(result)-1; i++ {
		// fmt.Println(len(result[i]))
		for j := 1; j < len(result[i])-1; j++ {
			if result[i][j] == "A" && i+1 < len(result[i+1]) {
				adj1 := result[i][j+1] + result[i][j-1]
				adj2 := result[i+1][j] + result[i-1][j]
				if strings.Count(adj1, "M") == 1 && strings.Count(adj1, "S") == 1 && strings.Count(adj2, "M") == 1 && strings.Count(adj2, "S") == 1 {
					count++
				}
			}
		}
	}
	return count
}

func combine_mas_string(matrix [][]string) string {
	combined := ""
	for index, val := range matrix {
		line := strings.Repeat(" ", int(math.Abs(float64(len(val)-(len(matrix[len(matrix)/2])))))/2) + strings.Join(val, "") + strings.Repeat(" ", int(math.Abs(float64(index-(len(matrix))/2))))
		//if len(matrix[len(matrix)/2]) == len(matrix[(len(matrix)/2)-1]) && index < len(matrix)/2 {
		//	line = " " + line
		//}
		combined = combined + line + "\n"
	}
	return combined
}

func count_occurances(corpus string, word string) int {
	lines := strings.Split(corpus, "\n")
	var result [][]string
	for _, line := range lines {
		words := strings.Split(line, "")
		result = append(result, words)
	}
	result = result[:len(result)-1]
	// fmt.Println(result)
	// fmt.Println(combine_string(result))
	// fmt.Println(combine_string(transpose(result)))
	// horizontal
	// forward_horiz := strings.Count(corpus, word)
	// // horizontal
	// backward_horiz := strings.Count(corpus, reverse(word))

	// // vertical

	// forward_vert := strings.Count(combine_string(transpose(result)), word)
	// // vertical
	// backward_vert := strings.Count(combine_string(transpose(result)), reverse(word))

	// diagnol
	// fmt.Println(combine_string(rotate(result)))

	// forward_diag := strings.Count(combine_string(rotate(result)), word)
	// diagnol
	// backward_diag := strings.Count(combine_string(rotate(result)), reverse(word))
	// fmt.Println(forward_horiz, backward_horiz, forward_vert, backward_vert, forward_diag, backward_diag)
	return -1 // forward_diag + backward_diag
}

func combine_string(matrix [][]string) string {
	combined := ""
	for _, val := range matrix {
		line := strings.Join(val, "")
		combined = combined + line + "\n"
	}
	return combined
}

// handle backwards
func reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

// handle vertical
func transpose(s [][]string) [][]string {
	var remade [][]string
	for i := 0; i < len(s[0]); i++ {
		var line []string
		for j := 0; j < len(s); j++ {
			line = append(line, s[j][i])
		}
		remade = append(remade, line)
	}
	return remade
}

// handle diagnol
func rotate(s [][]string) ([][]string, [][]string) {
	// fmt.Println(combine_string(s))
	// fmt.Println(s[9][9])
	var forward [][]string
	var backward [][]string
	x := len(s)
	y := len(s[0])
	// fmt.Println(x)
	for i := 0; i < y-int(math.Abs(float64(y-x))); i++ {
		var upper_l_line []string
		var lower_l_line []string
		var upper_r_line []string
		var lower_r_line []string
		for j := i; j >= 0; j-- {
			k := i - j
			upper_l_line = append(upper_l_line, s[j][k])
			lower_l_line = append(lower_l_line, s[x-j-1][y-k-1])
			upper_r_line = append(upper_r_line, s[x-j-1][k])
			lower_r_line = append(lower_r_line, s[j][y-k-1])
		}
		forward = append(forward, upper_l_line)
		backward = append(backward, lower_r_line)
		if i < y-1-int(math.Abs(float64(y-x))) {
			backward = append(backward, upper_r_line)
			forward = append(forward, lower_l_line)
		}
	}
	var final_l [][]string
	for i := 0; i < len(forward); i = i + 2 {
		final_l = append(final_l, forward[i])
	}
	for i := len(forward) - 2; i > 0; i = i - 2 {
		final_l = append(final_l, forward[i])
	}

	var final_r [][]string
	var final_outer [][]string
	for i := 1; i < len(final_l)-1; i = i + 2 {
		if i > len(final_l)/2 {
			final_r = append(final_r, strings.Split(reverse(strings.Join(final_l[i], "")), ""))
		} else {
			final_r = append(final_r, final_l[i])
		}
	}
	for i := 0; i < len(final_l); i = i + 2 {
		if i > len(final_l)/2 {
			final_outer = append(final_outer, strings.Split(reverse(strings.Join(final_l[i], "")), ""))
		} else {
			final_outer = append(final_outer, final_l[i])
		}
	}
	// fmt.Println(final_outer)
	// fmt.Println(final_r)
	return final_outer, final_r
}
