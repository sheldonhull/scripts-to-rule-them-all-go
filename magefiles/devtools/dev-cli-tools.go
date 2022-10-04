package devtools

// tools is a list of tools that are installed as binaries for development usage.
// This list gets installed to go bin directory once `mage init` is run.
// This is for binaries that need to be invoked as cli tools, not packages.
// Try to use aqua instead when possible to speed up builds and CI tooling, and use binaries when possible, not go install.
var toolList = []string{ //nolint:gochecknoglobals,deadcode,unused // ok to be global for tooling setup

}
