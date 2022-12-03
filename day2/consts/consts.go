package consts

var OPPONENT_CHOICES = map[string]int{
	"A": 1, // rock
	"B": 2, // paper
	"C": 3, // scissors
}

var MY_CHOICES = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

var WIN = 6
var DRAW = 3
var LOSE = 0

var RESULTS = [][]int{
	{DRAW, LOSE, WIN},
	{WIN, DRAW, LOSE},
	{LOSE, WIN, DRAW},
}

var LOSE_MAP = map[string]string{
	"A": "Z",
	"B": "X",
	"C": "Y",
}

var WIN_MAP = map[string]string{
	"A": "Y",
	"B": "Z",
	"C": "X",
}

var DRAW_MAP = map[string]string{
	"A": "X",
	"B": "Y",
	"C": "Z",
}

var EXPECTED_RESULTS = map[string]int{
	"X": LOSE, // lose
	"Y": DRAW, // draw
	"Z": WIN,  // win
}
