package constants

// Since we are dealing with builds, having a constants file until using a config input makes it easy.

// Default target to run when none is specified
// If not set, running mage will list available targets
// var Default = Build.

const (
	// ArtifactDirectory is a directory containing artifacts for the project and shouldn't be committed to source.
	ArtifactDirectory = ".artifacts"
	// CacheDirectory is a directory containing cached files that aren't specifically artifacts to keep discarding.
	// For example, include a development kubeconfig for a project, templating files, ssh keys or other files.
	CacheDirectory = ".cache"
)

const (
	// PermissionUserReadWriteExecute is the octal permission for read, write, & execute only for owner.
	PermissionUserReadWriteExecute = 0o0700
	// PermissionReadWriteSearchAll is the octal permission for all users to read, write, and search a file.
	PermissionReadWriteSearchAll = 0o0777
)

const (
	// defaultTrunkBranch is set to the upstream default and hard coded cause it shouldn't ever change again with this repo.
	DefaultTrunkBranch = "main"
)
