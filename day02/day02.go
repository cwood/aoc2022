package day02

const (
	Win  = 6
	Draw = 3
	Lost = 0

	TRock     = "A"
	TPaper    = "B"
	TScissors = "C"

	URock     = "X" // 1
	UPaper    = "Y" // 2
	UScissors = "Z" // 3

	ULose = "X"
	UDraw = "Y"
	UWin  = "Z"
)

func ScoreByResult(plays [][]string) int {
	var totalScore = 0

	for _, play := range plays {

		t := play[0]
		u := play[1]

		switch t {
		case TRock:
			switch u {
			case UWin:
				totalScore += Win + 2 // Paper
			case ULose:
				totalScore += Lost + 3 // Scissors
			case UDraw:
				totalScore += Draw + 1 // Rock
			}
		case TPaper:
			switch u {
			case UWin:
				totalScore += Win + 3 // Scissors
			case ULose:
				totalScore += Lost + 1 // Rock
			case UDraw:
				totalScore += Draw + 2 // Paper
			}
		case TScissors:
			switch u {
			case UWin:
				totalScore += Win + 1 // Rock
			case ULose:
				totalScore += Lost + 2 // Paper
			case UDraw:
				totalScore += Draw + 3 // Scissors
			}
		}
	}

	return totalScore
}

func ScoreByType(plays [][]string) int {

	var totalScore = 0

	for _, play := range plays {

		t := play[0]
		u := play[1]

		switch t {
		case TRock:
			switch u {
			case URock:
				totalScore += Draw + 1
			case UPaper:
				totalScore += Win + 2
			case UScissors:
				totalScore += Lost + 3
			}
		case TPaper:
			switch u {
			case UPaper:
				totalScore += Draw + 2
			case URock:
				totalScore += Lost + 1
			case UScissors:
				totalScore += Win + 3
			}
		case TScissors:
			switch u {
			case UScissors:
				totalScore += Draw + 3
			case UPaper:
				totalScore += Lost + 2
			case URock:
				totalScore += Win + 1
			}
		}
	}

	return totalScore
}
