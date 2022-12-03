package utils

import (
	"day2/consts"
)

func Compare(opponentChoice string, myChoice string) int {
	return consts.RESULTS[consts.MY_CHOICES[myChoice]-1][consts.OPPONENT_CHOICES[opponentChoice]-1]
}

func CalcResult(opponentChoice string, myChoice string) int {
	return Compare(opponentChoice, myChoice) + consts.MY_CHOICES[myChoice]
}
