package day08

func TreesVisibleOnEdge(in [][]int) int {
	var totalVisible = 0

	var edges = [][]int{
		[]int{0, 0},              // top 0
		[]int{0, len(in[0]) - 1}, // right 1
		[]int{0, 0},              // left 2
		[]int{len(in) - 1, 0},    // bottom 3
	}

	for e, pos := range edges {

		// mindex or move index is why x or y is being moved up
		var mindex = 0

		switch e {
		case 0:
		case 3:
			mindex = 0
		case 1:
		case 2:
			mindex = 1
		}
	}

	return totalVisible
}
