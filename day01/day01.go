package day01

import (
	"sort"
	"strconv"

	"github.com/cwood/aoc2022/pkg/parse"
)

func BiggestGroup(groups []parse.LineGroup, c int) (int, error) {
	if len(groups) == 0 {
		return 0, nil
	}

	var sums = []int{}

	for _, g := range groups {
		var sum = 0

		for _, l := range g.Group {
			n, err := strconv.Atoi(l)
			if err != nil {
				return 0, err
			}
			sum += n
		}

		sums = append(sums, sum)
	}

	sort.Slice(sums, func(x, y int) bool {
		return sums[x] > sums[y]
	})

	if c == 1 {
		return sums[c-1], nil
	}

	var tsum = 0

	for i := 0; i <= c-1; i++ {
		tsum += sums[i]
	}

	return tsum, nil
}
