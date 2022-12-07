package day05

import (
	"strings"
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

func TestTopStacksFromContainers(t *testing.T) {
	t.Run("Day 5 Part 1 Test Input", func(t *testing.T) {
		input, err := file.ReadLines("testinput")
		require.NoError(t, err)

		cntr, err := ParseInput(input)
		require.NoError(t, err)

		assert.Equal(t, []string{"C", "M", "Z"}, TopStacksFromContainer(cntr, true))
	})

	t.Run("Day 5 Part 1 Input", func(t *testing.T) {
		input, err := file.ReadLines("input")
		require.NoError(t, err)

		cntr, err := ParseInput(input)
		require.NoError(t, err)

		t.Logf("Day 5 Part 1: %s", strings.Join(TopStacksFromContainer(cntr, true), ""))
	})

	t.Run("Day 5 Part 2 Test Input", func(t *testing.T) {
		input, err := file.ReadLines("testinput")
		require.NoError(t, err)

		cntr, err := ParseInput(input)
		require.NoError(t, err)

		assert.Equal(t, []string{"M", "C", "D"}, TopStacksFromContainer(cntr, false))
	})

	t.Run("Day 5 Part 2 Input", func(t *testing.T) {
		input, err := file.ReadLines("input")
		require.NoError(t, err)

		cntr, err := ParseInput(input)
		require.NoError(t, err)

		t.Logf("Day 5 Part 2: %s", strings.Join(TopStacksFromContainer(cntr, false), ""))
	})
}
