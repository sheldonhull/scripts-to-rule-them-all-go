// ⚡ Core Mage Tasks
package main

import (
	"os"

	"github.com/magefile/mage/mg"
	"github.com/pterm/pterm"
	"github.com/sheldonhull/magetools/ci"
	"github.com/sheldonhull/magetools/fancy"
	"github.com/sheldonhull/magetools/tooling"


	// mage:import
	"github.com/sheldonhull/magetools/docgen"
	// mage:import
	"github.com/sheldonhull/magetools/gotools"
)

// 📁 createDirectories creates the local working directories for build artifacts and tooling.
func createDirectories() error {
	for _, dir := range []string{artifactDirectory,cacheDirectory} {
		if err := os.MkdirAll(dir, permissionUserReadWriteExecute); err != nil {
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

	pterm.DefaultSection.Println("CI Tooling")
	if err := tooling.SilentInstallTools(CIToolList); err != nil {
		return err
	}

	if ci.IsCI() {
		pterm.Success.Println("done with CI specific tooling. since detected in CI context, ending init early as core requirements met")
		return nil
	}

	mg.SerialDeps(
		(gotools.Go{}.Tidy),
		(gotools.Go{}.Init),
	)

	pterm.DefaultSection.Println("Setup Project Specific Tools")
	if err := tooling.SilentInstallTools(ToolList); err != nil {
		return err
	}
	return nil
}

// 🗑️ Clean up after yourself.
func Clean() {
	pterm.Success.Println("Cleaning...")
	for _, dir := range []string{artifactDirectory,cacheDirectory} {
		err := os.RemoveAll(dir)
		if err != nil {
			pterm.Error.Printf("failed to removeall: [%s] with error: %v\n", dir, err)
		}
		pterm.Success.Printf("🧹 [%s] dir removed\n", dir)
	}
	mg.Deps(createDirectories)
}
