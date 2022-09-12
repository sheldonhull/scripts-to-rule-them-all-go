//go:build mage

package main

// tools is a list of tools that are installed as binaries for development usage.
// This list gets installed to go bin directory once `mage init` is run.
// This is for binaries that need to be invoked as cli tools, not packages.
var ToolList = []string{ //nolint:gochecknoglobals // ok to be global for tooling setup
	"mvdan.cc/gofumpt@latest",
	"github.com/iwittkau/mage-select@latest",
	"github.com/mfridman/tparse@latest",     // Tparse provides nice formatted go test console output.

}
// CIToolList is the core tooling that `init` should run in the context of a CI system.
// This minimizes wasted cycles by focusing only on CI tooling and not developer experience.

var CIToolList = []string{ //nolint:gochecknoglobals // ok to be global for tooling setup
	"github.com/golangci/golangci-lint/cmd/golangci-lint@latest",
	"github.com/goreleaser/goreleaser@latest",
	"gotest.tools/gotestsum@latest",         // Gotestsum provides improved console output for tests as well as additional test output for CI systems.

}
