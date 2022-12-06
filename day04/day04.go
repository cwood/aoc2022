package day04

func Overlapping(input [][][]int) int {

	var overlapping = 0

	for i := 0; i < len(input); i += 1 {
		if input[i][0][0] <= input[i][1][0] && input[i][0][1] >= input[i][1][1] {
			overlapping += 1
		} else if input[i][1][0] <= input[i][0][0] && input[i][1][1] >= input[i][0][1] {
			overlapping += 1
		}
	}

	return overlapping
}
