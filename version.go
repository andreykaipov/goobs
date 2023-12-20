package goobs

import (
	"runtime/debug"
	"strings"
)

const lib = "github.com/andreykaipov/goobs"

var ProtocolVersion = "5.2.3"

var LibraryVersion = func() string {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		return "unknown"
	}

	for _, dep := range bi.Deps {
		if dep.Path == lib {
			return strings.TrimPrefix(dep.Version, "v")
		}
	}

	return "unknown"
}()
