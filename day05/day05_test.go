package day05

import (
	"testing"

	"github.com/cwood/aoc2022/pkg/file"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStackParsing(t *testing.T) {
	t.Run("Day 5 Part 1 Test Input", func(t *testing.T) {
		input, err := file.ReadLines("testinput")
		require.NoError(t, err)

		cntr, err := ParseInput(input)
		require.NoError(t, err)

		assert.Len(t, cntr.Stacks, 3)
		assert.Len(t, cntr.Stacks[2], 3)
		assert.Len(t, cntr.Instructions, 4)
		assert.Len(t, cntr.Instructions[0], 3)
	})

	t.Run("Day 5 Part 1 Input", func(t *testing.T) {
		input, err := file.ReadLines("input")
		require.NoError(t, err)

		cntr, err := ParseInput(input)
		require.NoError(t, err)

		assert.Len(t, cntr.Stacks, 9)
	})
}
