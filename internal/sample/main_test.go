package main

import (
	"runtime/debug"
	"testing"

	"github.com/andreykaipov/goobs"
	"github.com/stretchr/testify/assert"
)

const lib = "github.com/andreykaipov/goobs"

func TestVersion(t *testing.T) {
	assert.NotEmpty(t, goobs.LibraryVersion)
	assert.NotEmpty(t, goobs.ProtocolVersion)

	t.Run("build info", func(t *testing.T) {
		bi, ok := debug.ReadBuildInfo()
		assert.True(t, ok)

		actual := ""
		for _, dep := range bi.Deps {
			if dep.Path == lib {
				actual = dep.Version
			}
		}

		expected := "v" + goobs.LibraryVersion
		assert.NotEmpty(t, actual)
		assert.Equal(t, expected, actual)
	})
}
