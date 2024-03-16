// ⚡ Core Mage Tasks
package main

import (
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/pterm/pterm"
	"github.com/sheldonhull/magetools/ci"
	"github.com/sheldonhull/magetools/fancy"
	"github.com/sheldonhull/magetools/pkg/magetoolsutils"
	"github.com/sheldonhull/magetools/trunk"
	"template.local/magefiles/constants"

	// mage:import
	_ "github.com/sheldonhull/magetools/docgen"
	// mage:import
	"github.com/sheldonhull/magetools/gotools"
)

// createDirectories creates the local working directories for build artifacts and tooling.
func createDirectories() error {
	magetoolsutils.CheckPtermDebug()
	for _, dir := range []string{constants.ArtifactDirectory, constants.CacheDirectory} {
		if err := os.MkdirAll(dir, constants.PermissionUserReadWriteExecute); err != nil {
			pterm.Error.Printf("failed to create dir: [%s] with error: %v\n", dir, err)

			return err
		}
		pterm.Success.Printf("✅ [%s] dir created\n", dir)
	}

	return nil
}

// ⚡ Init runs multiple tasks to initialize all the requirements for running a project for a new contributor.
func Init() error {
	fancy.IntroScreen(ci.IsCI())
	pterm.Success.Println("running Init()...")

	mg.SerialDeps(
		Clean,
		createDirectories,
	)

	if ci.IsCI() {
		pterm.Success.Println(
			"done with CI specific tooling. since detected in CI context, ending init early as core requirements met",
		)
		return nil
	}

	mg.SerialDeps(
		(gotools.Go{}.Tidy),
		(gotools.Go{}.Init),
		(trunk.Trunk{}.Install),
	)
	// Aqua install is run in devcontainer/codespace automatically.
	// If this environment isn't being used, try to jump start, but if failure, output warning and let the developer choose if they want to go install or not.
	if err := sh.RunV("aqua", "install"); err != nil {
		pterm.Warning.Printfln("aqua install not successful.\n" +
			"This is optional, but will ensure every tool for the project is installed and matching version." +
			"To install see developer docs or go to https://aquaproj.github.io/docs/reference/install")
	}
	return nil
}

// 🗑️ Clean up after yourself.
func Clean() {
	pterm.Success.Println("Cleaning...")
	for _, dir := range []string{constants.ArtifactDirectory, constants.CacheDirectory} {
		err := os.RemoveAll(dir)
		if err != nil {
			pterm.Error.Printf("failed to removeall: [%s] with error: %v\n", dir, err)
		}
		pterm.Success.Printf("🧹 [%s] dir removed\n", dir)
	}
	mg.Deps(createDirectories)
}

// 🧹 DeepClean removes both artifacts and cache directory contents.
// Use this when you want to start over including any locally cached certs, files, or other things that normally you'd preserve between test runs.
func DeepClean() {
	magetoolsutils.CheckPtermDebug()
	pterm.Info.Println("🔥 Deep Cleaning...")
	for _, dir := range []string{constants.ArtifactDirectory, constants.CacheDirectory} {
		err := sh.Rm(dir)
		if err != nil {
			pterm.Error.Printf("failed to removeall: [%s] with error: %v\n", dir, err)
		}
		pterm.Success.Printf("🧹 [%s] dir removed\n", dir)
	}
	mg.Deps(createDirectories)
}
