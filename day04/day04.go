package day04

func ContainsAll(input [][][]int) int {

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

func ContainsAny(input [][][]int) int {
	var containsany = 0

	for _, grp := range input {
		if grp[1][0] <= grp[0][0] && grp[0][0] <= grp[1][1] {
			containsany += 1
		} else if grp[1][0] <= grp[0][1] && grp[0][1] <= grp[1][1] {
			containsany += 1
		} else if grp[0][0] <= grp[1][0] && grp[1][0] <= grp[0][1] {
			containsany += 1
		} else if grp[0][0] <= grp[1][1] && grp[1][1] <= grp[0][1] {
			containsany += 1
		}
	}

	return containsany
}
