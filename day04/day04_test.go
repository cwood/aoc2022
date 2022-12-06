package day04

import (
	"testing"

	"github.com/cwood/aoc2022/pkg/file"
	"github.com/cwood/aoc2022/pkg/parse"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOverlapping(t *testing.T) {
	t.Run("Day 4 Test Part 1", func(t *testing.T) {
		input, err := file.ReadLines("testinput")
		require.NoError(t, err)

		ranges, err := parse.ToRanges(input)
		require.NoError(t, err)

		assert.Equal(t, 2, Overlapping(ranges))
	})

	t.Run("Day 4 Part 1", func(t *testing.T) {
		input, err := file.ReadLines("input")
		require.NoError(t, err)

		ranges, err := parse.ToRanges(input)
		require.NoError(t, err)

		t.Logf("Day 4 Part 1: %d", Overlapping(ranges))
	})
}
