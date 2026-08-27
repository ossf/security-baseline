# Contributing to the Security Baseline

## Contributing to the Controls:

Currently, all the baseline controls are in `baseline.yaml`; follow the structure
[in the README](./README.md#baseline-structure) when proposing new entries.

## Contributing to the Tooling:

The baseline is published to https://baseline.openssf.org/ (via GitHub Pages) via
Jekyll (a static site generator) using scripts from `./cmd` and formatting from
`./docs`, using GitHub Actions.

## PR guidelines

All changes to the repository should be made via PR
([OSPS-AC-03](https://baseline.openssf.org/#osps-ac-03)).  In addition to a clear
title and descriptive commit message, PRs MUST meet the following criteria:

* DCO signoff (via `git commit -s` -- [OSPS-LE-01](https://baseline.openssf.org/#osps-le-01))
* All checks must pass ([OSPS-QA-04](https://baseline.openssf.org/#osps-qa-04))

### Check Go Tooling Linter

The OSPS Baseline tools are written in Go and the repository enforces linting on
every pull request. Before opening a PR, you can test your changes make the linter
happy by running [golangci-lint](https://golangci-lint.run/) locally in
the `cmd/` directory:

```bash
cd cmd/
golangci-lint run
```

### CSpell Check and Dictionary

The repo will enforce spell checks across the codebase. If you introduce new words
that the spell checker does not recognize, just add them to the `.cspell.yaml` file.

## Cutting a Release (maintainers)

1. Compile the versioned pages, pinning the mapping-document links to the release tag:
   `cd cmd && go run . compile --output ../docs/versions/YYYY-MM-DD.md --checklist-output ../docs/versions/YYYY-MM-DD-checklist.md --crosswalk-output ../docs/versions/YYYY-MM-DD-crosswalk.md --artifact-version vYYYY.MM.DD`
2. Hand-edit the compiled headers: replace `Version: devel` and the warning block with
   `Version: YYYY.MM.DD`, and add the `nav-title: Current Version` front matter to the
   main page (moving it off the previous release's page).
3. Update `docs/_config.yml` (`nav_pages`), the current/previous version lists in
   `docs/index.md`, and publish the drafted section in `docs/release_notes.md`.
4. Open the release PR; reviewers can preview the site with `make dev`.
5. After merging, create a tag **and a GitHub Release** for `vYYYY.MM.DD`
   (e.g. `gh release create vYYYY.MM.DD`). The Release — not the tag alone — triggers
   `publish.yaml`, which publishes the catalog and mapping documents to grc.store.
   Until it runs, the pinned grc.store links on the new version page will 404, and the
   site itself deploys as soon as the PR merges — so cut the Release promptly.

## Maintainer Status

See [./governance/GOVERNANCE.md](./governance/GOVERNANCE.md#maintainer-status) for
the process of achieving maintainer status on the project.
