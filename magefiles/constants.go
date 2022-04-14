//go:build mage

package main

// Since we are dealing with builds, having a constants file until using a config input makes it easy.

// Default target to run when none is specified
// If not set, running mage will list available targets
// var Default = Build.

// artifactDirectory is a directory containing artifacts for the project and shouldn't be committed to source.
const artifactDirectory = ".artifacts"

const permissionUserReadWriteExecute = 0o0700
