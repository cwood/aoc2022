package day03

import (
	"strings"

	"github.com/cwood/aoc2022/pkg/collections"
	"github.com/cwood/aoc2022/pkg/parse"
)

var (
	priorities = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func SumGroups(grps []parse.LineGroup) int {
	var totalSum = 0

	for _, grp := range grps {

		var counter = collections.Counter{}

		for _, l := range grp.Group {
			s := collections.Set{}
			for _, r := range l {
				s.Add(string(r))
			}

			for k := range s {
				counter.Add(k)
			}
		}

		for v, c := range counter {
			if c == len(grps[0].Group) { // Make the assumption we need to match all N groups
				if v := strings.Index(priorities, v.(string)); v != -1 {
					totalSum += v + 1
				}
			}
		}

	}

	return totalSum
}

func SumPriorities(input []string) int {

	var totalSum = 0

inputs:
	for _, l := range input {
		half := len(l) / 2
		startStr, endStr := l[:half], l[half:]

		start := collections.Set{}

		for _, s := range startStr {
			start.Add(s)
		}

		for _, e := range endStr {
			if start.Contains(e) {
				if v := strings.Index(priorities, string(e)); v != -1 {
					totalSum += v + 1
				}

				continue inputs
			}
		}
	}

	return totalSum
}
