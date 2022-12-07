package day06

import (
	"log"

	"github.com/cwood/aoc2022/pkg/collections"
)

func FindMarkerInLine(in string) int {
	if len(in) < 4 {
		return 0
	}

	for i := 4; i <= len(in)-1; i++ {
		window := in[i-4 : i]

		c := collections.Counter{}
		for _, r := range window {
			c.Add(r)
		}

		if c.All(1) {
			log.Print(i, c, window)
			return i
		}
	}

	return 0
}
