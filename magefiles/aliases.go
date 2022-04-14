//go:build mage

package main

var Aliases = map[string]interface{}{ //nolint:deadcode,gochecknoglobals // it's ok for this to be sitting out here as it's a mage feature
	"bootstrap": Init,
}
