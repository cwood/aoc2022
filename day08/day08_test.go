package day08

import (
	"testing"

	"github.com/cwood/aoc2022/pkg/file"
	"github.com/cwood/aoc2022/pkg/parse"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTreesVisibileOnEdge(t *testing.T) {
	t.Run("Day 8 Test Part 1", func(t *testing.T) {
		input, err := file.ReadLines("testinput")
		require.NoError(t, err)

		grid, err := parse.ToGrid(input)
		require.NoError(t, err)

		assert.Equal(t, 21, TreesVisibleOnEdge(grid))
	})
}
