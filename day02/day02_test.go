package day02

import (
	"testing"

	"github.com/cwood/aoc2022/pkg/file"
	"github.com/cwood/aoc2022/pkg/parse"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTotalScore(t *testing.T) {
	t.Run("Day 2 Test Problem 1", func(t *testing.T) {
		testInput, err := file.ReadLines("testinput")
		require.NoError(t, err)

		plays := parse.ToWords(testInput)

		require.Len(t, plays, 3)
		require.Len(t, plays[0], 2)

		assert.Equal(t, ScoreByType(plays), 15)
	})

	t.Run("Day 2 Problem 1", func(t *testing.T) {
		testInput, err := file.ReadLines("input")
		require.NoError(t, err)

		plays := parse.ToWords(testInput)
		t.Logf("Day 2 Problem 1: %d", ScoreByType(plays))
	})

	t.Run("Day 2 Test Problem 2", func(t *testing.T) {
		testInput, err := file.ReadLines("testinput")
		require.NoError(t, err)

		plays := parse.ToWords(testInput)

		require.Len(t, plays, 3)
		require.Len(t, plays[0], 2)

		assert.Equal(t, ScoreByResult(plays), 12)
	})

	t.Run("Day 2 Problem 2", func(t *testing.T) {
		testInput, err := file.ReadLines("input")
		require.NoError(t, err)

		plays := parse.ToWords(testInput)
		t.Logf("Day 2 Problem 2: %d", ScoreByResult(plays))
	})

}
