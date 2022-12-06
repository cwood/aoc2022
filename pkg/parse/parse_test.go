package parse

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestToLineGroup(t *testing.T) {
	g := ToLineGroupWithSep([]string{
		"a",
		"b",
		"c",
		"",
		"a",
		"b",
		"",
		"a",
	})

	assert.Len(t, g, 3)
}

func TestToRanges(t *testing.T) {
	grps, err := ToRanges([]string{
		"1-2,3-5,4-5",
		"2-3,6-7",
	})
	require.NoError(t, err)

	assert.Len(t, grps, 2)
	assert.Len(t, grps[0], 3)
	assert.Len(t, grps[0][1], 2)
	assert.Equal(t, grps[0][0][0], 1)
	assert.Equal(t, grps[1][1][1], 7)
}
