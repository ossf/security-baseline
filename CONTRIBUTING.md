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
title and descriptive commit message, PRs MUST meet the following control:

* DCO signoff (via `git commit -s` -- [OSPS-LE-01](https://baseline.openssf.org/#osps-le-01))
* All checks must pass ([OSPS-QA-04](https://baseline.openssf.org/#osps-qa-04))

## Maintainer Status

See [./governance/GOVERNANCE.md](./governance/GOVERNANCE.md#maintainer-status) for
the process of achieving maintainer status on the project.