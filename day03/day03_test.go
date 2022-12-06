package day03

import (
	"testing"

	"github.com/cwood/aoc2022/pkg/file"
	"github.com/cwood/aoc2022/pkg/parse"
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

	t.Run("Day 3 Test Problem 2", func(t *testing.T) {
		testInput, err := file.ReadLines("testinput")
		require.NoError(t, err)

		lgrps := parse.ToLineGroupWithCnt(testInput, 3)
		require.Len(t, lgrps, 2)
		require.Len(t, lgrps[0].Group, 3)
		require.Len(t, lgrps[1].Group, 3)

		assert.Equal(t, 70, SumGroups(lgrps))
	})

	t.Run("Day 3 Problem 2", func(t *testing.T) {
		testInput, err := file.ReadLines("input")
		require.NoError(t, err)

		lgrps := parse.ToLineGroupWithCnt(testInput, 3)

		t.Logf("Day 3 Problem 2: %d", SumGroups(lgrps))
	})
}
