package parse

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
