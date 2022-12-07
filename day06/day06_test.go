package day06

import (
	"testing"

	"github.com/cwood/aoc2022/pkg/file"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFindMarkerInLine(t *testing.T) {
	t.Run("Day 6 Part 1 Test Inputs", func(t *testing.T) {
		assert.Equal(t, 7, FindMarkerInLine("mjqjpqmgbljsphdztnvjfqwrcgsmlb"))
		assert.Equal(t, 5, FindMarkerInLine("bvwbjplbgvbhsrlpgdmjqwftvncz"))
		assert.Equal(t, 6, FindMarkerInLine("nppdvjthqldpwncqszvftbrmjlhg"))
		assert.Equal(t, 10, FindMarkerInLine("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"))
		assert.Equal(t, 11, FindMarkerInLine("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"))
	})

	t.Run("Day 6 Part 1", func(t *testing.T) {
		input, err := file.ReadLines("input")
		require.NoError(t, err)

		t.Logf("Day 6 Part 1: %d", FindMarkerInLine(input[0]))
	})

	t.Run("Day 6 Part 2 Test Inputs", func(t *testing.T) {
		assert.Equal(t, 19, FindMarkerInLine("mjqjpqmgbljsphdztnvjfqwrcgsmlb", WithWindowLength(14)))
		assert.Equal(t, 23, FindMarkerInLine("bvwbjplbgvbhsrlpgdmjqwftvncz", WithWindowLength(14)))
		assert.Equal(t, 23, FindMarkerInLine("nppdvjthqldpwncqszvftbrmjlhg", WithWindowLength(14)))
		assert.Equal(t, 29, FindMarkerInLine("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", WithWindowLength(14)))
		assert.Equal(t, 26, FindMarkerInLine("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", WithWindowLength(14)))
	})

	t.Run("Day 6 Part 2", func(t *testing.T) {
		input, err := file.ReadLines("input")
		require.NoError(t, err)

		t.Logf("Day 6 Part 2: %d", FindMarkerInLine(input[0], WithWindowLength(14)))
	})
}
