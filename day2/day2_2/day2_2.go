package main

import (
	"bufio"
	"day2/consts"
	"day2/utils"
	"fmt"
	"log"
	"os"
	"strings"
)

func resultToChoice(opponentChoice string, expectedResult string) string {
	if consts.EXPECTED_RESULTS[expectedResult] == consts.DRAW {
		return consts.DRAW_MAP[opponentChoice]
	}
	if consts.EXPECTED_RESULTS[expectedResult] == consts.LOSE {
		return consts.LOSE_MAP[opponentChoice]
	}
	if consts.EXPECTED_RESULTS[expectedResult] == consts.WIN {
		return consts.WIN_MAP[opponentChoice]
	}
	return ""
}

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
		args := strings.Fields(line)

		opponentChoice := args[0]
		expectedResult := args[1]

		myChoice := resultToChoice(opponentChoice, expectedResult)
		res := utils.CalcResult(opponentChoice, myChoice)
		sum += res

	}

	fmt.Println("Total: ", sum)
}
