package day07

import (
	"testing"

	"github.com/cwood/aoc2022/pkg/file"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewFileSystemFromInput(t *testing.T) {
	input, err := file.ReadLines("testinput")
	require.NoError(t, err)

	fs := NewFilesystemFromInput(input)
	require.NotNil(t, fs)

	assert.Len(t, fs.Directories, 2)
	assert.Len(t, fs.Files, 2)
	assert.Equal(t, fs.Files[0].Name, "b.txt")
	assert.Equal(t, fs.Directories[0].Directories[0].Files[0].Name, "i")
	assert.Equal(t, fs.Directories[0].Directories[0].Files[0].Size, int64(584))

	input, err = file.ReadLines("input")
	require.NoError(t, err)

	fs = NewFilesystemFromInput(input)
	require.NotNil(t, fs)
}

func TestDirectoriesNearist(t *testing.T) {
	t.Run("Day 7 Test Part 1", func(t *testing.T) {
		input, err := file.ReadLines("testinput")
		require.NoError(t, err)

		fs := NewFilesystemFromInput(input)
		require.NotNil(t, fs)

		assert.Equal(t, int64(95437), DirectoriesNeariest(fs, 100000))
	})

	t.Run("Day 7 Part 1", func(t *testing.T) {
		input, err := file.ReadLines("input")
		require.NoError(t, err)

		fs := NewFilesystemFromInput(input)
		require.NotNil(t, fs)

		t.Logf("Day 7 Part 1: %d", DirectoriesNeariest(fs, 100000))
	})
}
