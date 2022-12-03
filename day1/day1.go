package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var temp_sum_value int = 0
	var sums []int

	file, err := os.Open("day1.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := strings.TrimSuffix(scanner.Text(), "\n")

		if line == "" {
			sums = append(sums, temp_sum_value)
			temp_sum_value = 0
		} else {
			value, _ := strconv.Atoi(line)
			temp_sum_value += value
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sums)))
	fmt.Println("First three: ", sums[0], sums[1], sums[2])
	fmt.Println("Total: ", sums[0]+sums[1]+sums[2])

}
