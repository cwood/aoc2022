package day03

import (
	"strings"

	"github.com/cwood/aoc2022/pkg/collections"
)

var (
	priorities = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

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
