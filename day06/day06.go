package day06

import (
	"github.com/cwood/aoc2022/pkg/collections"
)

type Cfg struct {
	windowLen int
}

type Option func(*Cfg)

func WithWindowLength(i int) Option {
	return func(c *Cfg) {
		c.windowLen = i
	}
}

func FindMarkerInLine(in string, opts ...Option) int {

	cfg := &Cfg{windowLen: 4}

	for _, opt := range opts {
		opt(cfg)
	}

	if len(in) < cfg.windowLen {
		return 0
	}

	for i := cfg.windowLen; i <= len(in)-1; i++ {
		window := in[i-cfg.windowLen : i]

		c := collections.Counter{}
		for _, r := range window {
			c.Add(r)
		}

		if c.All(1) {
			return i
		}
	}

	return 0
}
