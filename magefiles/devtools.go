//go:build mage

package main

// tools is a list of tools that are installed as binaries for development usage.
// This list gets installed to go bin directory once `mage init` is run.
// This is for binaries that need to be invoked as cli tools, not packages.
var toolList = []string{ //nolint:gochecknoglobals // ok to be global for tooling setup
	"github.com/golangci/golangci-lint/cmd/golangci-lint@latest",
	"github.com/dustinkirkland/golang-petname/cmd/petname@latest",
	"mvdan.cc/gofumpt@latest",
	"github.com/daixiang0/gci@latest",
	"github.com/goreleaser/goreleaser@latest",
	"github.com/iwittkau/mage-select@latest",
	"github.com/asiermarques/adrgen@latest",     // For managing ADR files and creating from template.
	"github.com/mfridman/tparse@latest",         // Tparse provides nice formatted go test console output.
	"github.com/rakyll/gotest@latest",           // Gotest is a wrapper for running Go tests via command line with support for colors to make it more readable.
	"gotest.tools/gotestsum@latest",             // Gotestsum provides improved console output for tests as well as additional test output for CI systems.
	"github.com/zricethezav/gitleaks/v8@latest", // Doesn't work right now due to replace directives.
}
