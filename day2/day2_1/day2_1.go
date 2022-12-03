package main

import (
	"bufio"
	"day2/utils"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	sum := 0

	file, err := os.Open("../day2.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := strings.TrimSuffix(scanner.Text(), "\n")
		choices := strings.Fields(line)

		opponentChoice := choices[0]
		myChoice := choices[1]

		res := utils.CalcResult(opponentChoice, myChoice)
		sum += res
	}

	fmt.Println("Total: ", sum)
}
