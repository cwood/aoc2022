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
