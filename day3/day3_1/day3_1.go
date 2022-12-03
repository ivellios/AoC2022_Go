package main

import (
	"bufio"
	"day3/utils"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/juliangruber/go-intersect"
	"github.com/samber/lo"
)

func rucksackReorganization(rucksack string) int {
	left := utils.StringToSet(rucksack[:len(rucksack)/2])
	right := utils.StringToSet(rucksack[len(rucksack)/2:])
	common := intersect.Simple(lo.Keys(left), lo.Keys(right))
	r := []rune(fmt.Sprintf("%v", common[0]))
	return utils.RuneToValue(r[0])
}

func main() {
	total := 0
	file, err := os.Open("../day3.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := strings.TrimSuffix(scanner.Text(), "\n")
		total += rucksackReorganization(line)
	}

	fmt.Println("Total: ", total)
}
