# scripts-to-rule-them-all-go

Inspired from the 2015 repo for "Scripts to Rule Them All", this applies that same concept but with using a more modern toolchain like Go and Mage.
Mage provides a Go based task runner that can replace makefiles, custom bash/python/powershell scripts, and instead provide a cross-platform common set of commands to run a project.

> 2022-03 🚧 WIP - Bringing in a few more examples. It's ready to use for basics though!

## Why Mage

- Mage is cross-platform.
- Mage is Go. Go is fun. Go helps write really robust safe code with the error handling.
- Easy to plug-in to CI since Mage is just Go.

## Why This Template?

This project is a template to help bootstrap a new project with a few things that normally result in more setup work.
I've tended to standardize on a few key things.

- If it's complicated and I do periodically, wrap up the bash script commands in Go instead if it's only a few minutes of work.
- If I can, I'll use a package, but for things that are more complicated (like Kubernetes), start with just wrapping up kubectl commands and go from there.

Normal tasks:

### Included & Prebaked

- `init`: All projects bootstrap from this. It runs go mod tidy, go installs, and is extended for any other tools as well.
  For example, I'll add `asdf:install` to the `init` task and let it also ensure all apps I want are setup.
- `devcontainer:build`: Locally run the steps to grab the image and build a local containerized devcontainer to work in.

## You Want More?

👉 Checkout [magetools](https://github.com/sheldonhull/magetools)

I've been steadily baking automation tasks that are tested and reusable in there.
Most of the tasks self-setup any tooling as well.

### Custom

- `publish`: Focused on CI based publishing with a tool like `goreleaser` or `ko`.
- `bump`: Semantic versions
- `doctor`: Anytime a setup/problem in the project is experienced I try to add a diagnostic check on versions of apps, env vars, and other things and return this as a pass/fail or info table to summarize in a very clean way. This uses pterm.

## Example Output

Running `mage` in a brand new project.

```text
⚡ Core Mage Tasks

Targets:
  clean              up after yourself.
  go:doctor          🏥 Doctor will provide config details.
  go:fmt             ✨ Fmt runs gofumpt.
  go:init            ⚙️ Init runs all required steps to use this package.
  go:lint            🔎 Run golangci-lint without fixing.
  go:lintConfig      🏥 LintConfig will return output of golangci-lint config.
  go:test            🧪 Run go test.
  go:testSum         🧪 Run gotestsum (Params: Path just like you pass to go test, ie ./..., pkg/, etc ).
  go:tidy            🧹 Tidy tidies.
  go:wrap            ✨ Wrap runs golines powered by gofumpt.
  init               runs multiple tasks to initialize all the requirements for running a project for a new contributor.
  secrets:detect     🔐 Detect scans for secret violations with gitleaks without git consideration.
  secrets:protect    🔐 Protect scans the staged artifacts for violations.
```

## Template

You can use this as a template to get started with a project ready to run using Mage.

## Setup

- Multiple options to install on [Mage - Docs](https://magefile.org/)
- Additional options:

### Using Go (won't have version info embedded)

```go
go install github.com/magefile/mage@latest
```

### Zinit

In your `.zshrc` or equivalent for [zinit](https://github.com/zdharma-continuum/zinit) plugins:

Using sbin for shimming:

```shell
zinit ice as"null" lucid from"github-rel" \
    mv"mage* -> mage"  \
    extract'' \
    lucid
zinit load magefile/mage
```

Without sbin:

```shell
zinit ice as"null" lucid from"github-rel" \
    mv"mage* -> mage"  \
    extract'' \
    lucid  \
    sbin'mage'
zinit load magefile/mage
```

## Note

Maintain `//go:build mage` to avoid impacting test coverage reports.

I've added this to the tags for language server build to recognize so you can use code completion.
