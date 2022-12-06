package day01

import (
	"testing"

	"github.com/cwood/aoc2022/pkg/file"
	"github.com/cwood/aoc2022/pkg/parse"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMostCaloriesByElf(t *testing.T) {
	t.Run("Day 1 Test Problem 1", func(t *testing.T) {
		testInput, err := file.ReadLines("testinput")
		require.NoError(t, err)

		lgs := parse.ToLineGroupWithSep(testInput)
		assert.Len(t, lgs, 5)

		g, err := BiggestGroup(lgs, 1)
		require.NoError(t, err)

		assert.Equal(t, g, 24000)
	})

	t.Run("Day 1 Problem 1", func(t *testing.T) {
		input, err := file.ReadLines("input")
		require.NoError(t, err)

		lgs := parse.ToLineGroupWithSep(input)
		assert.NotZero(t, lgs)

		g, err := BiggestGroup(lgs, 1)
		require.NoError(t, err)

		t.Logf("Day 1 Problem 1: %d", g)
	})

	t.Run("Day 1 Test Problem 2", func(t *testing.T) {
		testInput, err := file.ReadLines("testinput")
		require.NoError(t, err)

		lgs := parse.ToLineGroupWithSep(testInput)
		assert.Len(t, lgs, 5)

		g, err := BiggestGroup(lgs, 3)
		require.NoError(t, err)

		assert.Equal(t, g, 45000)
	})

	t.Run("Day 1 Problem 2", func(t *testing.T) {
		input, err := file.ReadLines("input")
		require.NoError(t, err)

		lgs := parse.ToLineGroupWithSep(input)
		assert.NotZero(t, lgs)

		g, err := BiggestGroup(lgs, 3)
		require.NoError(t, err)

		t.Logf("Day 1 Problem 2: %d", g)
	})
}
