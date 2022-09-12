//go:build mage

package main

// Since we are dealing with builds, having a constants file until using a config input makes it easy.

// Default target to run when none is specified
// If not set, running mage will list available targets
// var Default = Build.

// artifactDirectory is a directory containing artifacts for the project and shouldn't be committed to source.
const artifactDirectory = ".artifacts"

// PermissionUserReadWriteExecute is the octal permission for read, write, & execute only for owner.
permissionUserReadWriteExecute = 0o0700

// permissionReadWriteSearchAll is the octal permission for all users to read, write, and search a file.
permissionReadWriteSearchAll = 0o0777
