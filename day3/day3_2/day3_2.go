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

func findBadge(group []string) rune {
	keys1 := lo.Keys(utils.StringToSet(group[0]))
	keys2 := lo.Keys(utils.StringToSet(group[1]))
	common1 := intersect.Simple(keys1, keys2)
	common := intersect.Simple(common1, lo.Keys(utils.StringToSet(group[2])))
	r := []rune(fmt.Sprintf("%v", common[0]))
	return r[0]
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

		line1 := strings.TrimSuffix(scanner.Text(), "\n")
		scanner.Scan()
		line2 := strings.TrimSuffix(scanner.Text(), "\n")
		scanner.Scan()
		line3 := strings.TrimSuffix(scanner.Text(), "\n")

		lines := []string{line1, line2, line3}

		total += utils.RuneToValue(findBadge(lines))
	}

	fmt.Println("Total: ", total)
}
