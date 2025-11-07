---
nav-title: Maintainer Guidance
---

# Implementation Guidance for Maintainers

Are you a project maintainer looking to evaluate your project against the OSPS Baseline?
This page contains guidance to help you do it quickly.
Your feedback is always welcome in the [issues](https://github.com/ossf/security-baseline/issues).

The Baseline website includes a Markdown-formatted checklist that you can use in a repository issue to track the controls by level.
See the [main page](/) for links to the checklist.

## Use Security Insights

The [Security Insights project](https://github.com/ossf/security-insights) maintains a specification for projects to report information about their security in a machine-processable way.
Many tools for Baseline evaluation will use this to evaluate controls that are not easily machine-auditable via platform APIs.

To get started with Security Insights:

. Adopt the Security Insights specification by creating a security-insights.yml file in your repository.
. Populate the security-insights.yml file with security data for your project following the example-full.yml template.

## Evaluation tooling

Several tools are available to help you automate the process.

* [LFX Insights](https://insights.linuxfoundation.org/) provides automated measurement and reporting of a variety of project metrics, including Baseline controls.
If your project is not already included in LFX Insights, you can [add it](https://github.com/linuxfoundation/insights/discussions/categories/project-onboardings?discussions_q=is:open+category:%22Project+onboardings%22+sort:top).
Your project does not need to be a part of the Linux Foundation.
* [Privateer Plugin for GitHub Repositories](https://github.com/revanite-io/pvtr-github-repo) performs automated checks of some Baseline controls.
This is what powers the LFX Insights evaluation.
It is also available as a [GitHub Action](https://github.com/revanite-io/pvtr-github-repo-action).

## Additional information

* OpenSSF case study for how [GUAC used LFX Insights](https://openssf.org/blog/2025/08/14/case-study-how-lfx-insights-and-osps-baseline-validated-guacs-security-in-under-an-hour/) to quickly evaluate against levels 1 and 2.