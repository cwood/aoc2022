package day03

import (
	"testing"

	"github.com/cwood/aoc2022/pkg/file"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSumPriorities(t *testing.T) {
	t.Run("Day 3 Test Problem 1", func(t *testing.T) {
		testInput, err := file.ReadLines("testinput")
		require.NoError(t, err)

		assert.Equal(t, 157, SumPriorities(testInput))
	})

	t.Run("Day 3 Problem 1", func(t *testing.T) {
		testInput, err := file.ReadLines("input")
		require.NoError(t, err)

		t.Logf("Day 3 Problem 1: %d", SumPriorities(testInput))
	})
}
