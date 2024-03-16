package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"template.local/magefiles/constants"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/pterm/pterm"
	"github.com/sheldonhull/magetools/pkg/magetoolsutils"
)

// Changelog contains task that batch up the changelog commands and allow triggering a release with an explicit version.
type Changelog mg.Namespace

// getVersion returns the version and path for the changefile to use for the semver and release notes.
func getVersion() (releaseVersion, cleanPath string, err error) {
	magetoolsutils.CheckPtermDebug()
	releaseVersion, err = sh.Output("changie", "latest")
	if err != nil {
		pterm.Error.Printfln("changie pulling latest release note version failure: %v", err)
		return "", "", err
	}
	cleanVersion := strings.TrimSpace(releaseVersion)
	cleanPath = filepath.Join(".changes", cleanVersion+".md")
	if os.Getenv("GITHUB_WORKSPACE") != "" {
		cleanPath = filepath.Join(os.Getenv("GITHUB_WORKSPACE"), ".changes", cleanVersion+".md")
	}
	return cleanVersion, cleanPath, nil
}

// 📦 Bump the application as an interactive command, prompting for semver change type, merging changelog, and running format and git add.
func (Changelog) Bump() error {
	magetoolsutils.CheckPtermDebug()
	pterm.DefaultSection.Println("(Changelog) Bump()")
	bumpType, _ := pterm.DefaultInteractiveSelect.
		WithOptions([]string{"auto", "major"}).
		Show()
	pterm.Info.Printfln("bumping by: %s", bumpType)
	if bumpType == "major" {
		pterm.Warning.Printfln(
			"🔥major bumping should be done with care as this signifies large breaking changes",
		)
		pterm.Warning.Println("making sure you grok this 😀")
		bumpType, _ = pterm.DefaultInteractiveSelect.
			WithOptions([]string{"patch", "minor", "major"}).
			Show()
	}
	if err := sh.RunV("changie", "batch", bumpType); err != nil {
		pterm.Warning.Printf(
			"changie batch failure (non-terminating as might be repeating batch command): %v",
			err,
		)
	}
	if err := sh.RunV("changie", "merge"); err != nil {
		return err
	}
	if err := sh.RunV("trunk", "fmt"); err != nil {
		return err
	}
	if err := sh.RunV("trunk", "check", "--ci"); err != nil {
		pterm.Warning.Printfln(
			"trunk check failure. This is non-terminating for the mage task, but you should check it before merging",
		)
	}
	if err := sh.RunV("git", "add", ".changes/*"); err != nil {
		return err
	}
	if err := sh.RunV("git", "add", "*.yaml"); err != nil {
		return err
	}
	if err := sh.RunV("git", "add", "*.md"); err != nil {
		return err
	}

	releaseVersion, cleanPath, err := getVersion()
	if err != nil {
		return err
	}
	pterm.Info.Println(" Are you ready to create a commit with these changes?")
	confirm, err := pterm.DefaultInteractiveConfirm.
		WithDefaultValue(false).
		WithRejectText("no").
		WithConfirmText("yes").
		Show()
	if err != nil {
		return err
	}
	if !confirm {
		pterm.Warning.Println("someone changed their mind")
		return nil
	}
	if err := sh.RunV("git", "commit", "-m", fmt.Sprintf("feat: 🚀 create release %s", releaseVersion)); err != nil {
		return err
	}
	pterm.Info.Println("Ready to tag and push?")
	confirm, err = pterm.DefaultInteractiveConfirm.
		WithDefaultValue(false).
		WithRejectText("no").
		WithConfirmText("yes").
		Show()
	if err != nil {
		return err
	}

	if confirm {
		if err := sh.RunV("git", "tag", "-a", releaseVersion, "--file", cleanPath); err != nil {
			return err
		}
		pterm.Success.Printfln("tag %s created", releaseVersion)

		if err := sh.RunV("git", "push", "--follow-tags"); err != nil {
			return err
		}
		pterm.Success.Printfln("tag %s pushed", releaseVersion)
	}
	return nil
}

// 🏷 Tag will use changie versioning to tag the git repo and should only be run on main.
//
// This runs the tagging, makes sure is set to the default upstream branch, and then pushes the tag.
// This is meant to be run by CI eventually, so that CI confirms tests pass then tags, and this tag is used to launch the release.
func (Changelog) Tag() error {
	magetoolsutils.CheckPtermDebug()

	pterm.DefaultSection.Println("(Changelog) Tag()")

	branch, err := sh.Output("git", "current-branch")
	if err != nil {
		return fmt.Errorf("getting current branch: %w", err)
	}
	pterm.Success.Printfln("current branch: %s", branch)
	if branch != constants.DefaultTrunkBranch {
		return fmt.Errorf(
			"changie tag only works with [%s] branch and this was run against: [%s] (make sure not a detached checkout)",
			constants.DefaultTrunkBranch,
			branch,
		)
	}

	releaseVersion, cleanPath, err := getVersion()
	if err != nil {
		return err
	}

	pterm.Info.Println("Ready to tag and push?")
	var confirm bool
	confirm, err = pterm.DefaultInteractiveConfirm.
		WithDefaultValue(false).
		WithRejectText("no").
		WithConfirmText("yes").
		Show()
	if err != nil {
		return err
	}

	if confirm {
		if err := sh.RunV("git", "tag", "-a", releaseVersion, "--file", cleanPath); err != nil {
			return err
		}
		pterm.Success.Printfln("tag %s created", releaseVersion)

		if err := sh.RunV("git", "push", "--follow-tags"); err != nil {
			return err
		}
		pterm.Success.Printfln("tag %s pushed", releaseVersion)
	}
	return nil
}

// 📦 Merge updates the changelog without bumping the version.
// This is useful for when you are picking up after the changie batch has already completed, but need to re-run the changie merge.
func (Changelog) Merge() error {
	magetoolsutils.CheckPtermDebug()
	pterm.DefaultSection.Println("(Changelog) Merge()")
	if err := sh.RunV("changie", "merge"); err != nil {
		return err
	}
	if err := sh.RunV("trunk", "fmt"); err != nil {
		return err
	}
	if err := sh.RunV("trunk", "check", "--ci"); err != nil {
		pterm.Warning.Printfln(
			"trunk check failure. This is non-terminating for the mage task, but you should check it before merging",
		)
	}

	if err := sh.RunV("git", "add", ".changes/*"); err != nil {
		return err
	}
	if err := sh.RunV("git", "add", "*.yaml"); err != nil {
		return err
	}
	if err := sh.RunV("git", "add", "*.md"); err != nil {
		return err
	}

	releaseVersion, _, err := getVersion()
	if err != nil {
		return err
	}
	pterm.Info.Println(" Are you ready to create a commit with these changes?")
	confirm, err := pterm.DefaultInteractiveConfirm.
		WithDefaultValue(false).
		WithRejectText("no").
		WithConfirmText("yes").
		Show()
	if err != nil {
		return err
	}
	if !confirm {
		pterm.Warning.Println("someone changed their mind")
		return nil
	}
	if err := sh.RunV("git", "commit", "-m", fmt.Sprintf("feat: 🚀 create release %s", releaseVersion)); err != nil {
		return err
	}

	return nil
}
