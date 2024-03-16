package main

// Since we are dealing with builds, having a constants file until using a config input makes it easy.

// Default target to run when none is specified
// If not set, running mage will list available targets
// var Default = Build.

const (
	// artifactDirectory is a directory containing artifacts for the project and shouldn't be committed to source.
	artifactDirectory = ".artifacts"
	// cacheDirectory is a directory containing cached files that aren't specifically artifacts to keep discarding.
	// For example, include a development kubeconfig for a project, templating files, ssh keys or other files.
	cacheDirectory = ".cache"
)

const (
	// PermissionUserReadWriteExecute is the octal permission for read, write, & execute only for owner.
	permissionUserReadWriteExecute = 0o0700
	// permissionReadWriteSearchAll is the octal permission for all users to read, write, and search a file.
	permissionReadWriteSearchAll = 0o0777
)
