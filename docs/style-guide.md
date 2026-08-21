# Open Source Project Security Baseline (OSPS) Documentation Style Guide

## About this guide

This guide is the official style guide for Open Source Project Security Baseline documentation. It defines the conventions for writing and maintaining OSPS Baseline controls.

The goal of the OSPS Documentation Style Guide is to ensure that every control mentioned in the Open Source Project Security Baseline is **clear, easy to maintain, unambiguous, testable and consistent** across all the specifications.

## Additional editorial resources

Here is a list of supplementary resources to research and help you think, but don’t consider them a part of the OSPS Documentation Style Guide:

- [Red Hat supplementary style guide for product documentation](https://redhat-documentation.github.io/supplementary-style-guide/)  
- [Google developer documentation style guide](https://developers.google.com/style)   
- [1Password Style Guide](https://support.1password.com/style-guide/)  
- [The New Oxford American Dictionary](https://www.oxfordreference.com/display/10.1093/acref/9780195392883.001.0001/acref-9780195392883) 



# Fundamentals

## Section Structure

Every control should follow the same order to simplify maintenance. The sections should follow the order specified in the security baseline. Refer to the `OSPS-${ControlFamilyAbbreviated}.yml` files for examples.

Do not rearrange sections in the YAML. If fields are listed in a different order from one entry to the next, it becomes difficult to see what content is complete or missing.

The two levels of identifiers mentioned in the baseline document are defined as follows:

| Format | Represents | Example |
|---|---|---|
| `OSPS-XX-YY` | A control (an objective and its constituent parts) | `OSPS-AC-03` |
| `OSPS-XX-YY.ZZ` | An individual requirement within that control | `OSPS-AC-03.01` |

## Stable Requirement Identifiers

Control and Requirement identifiers are permanent. Once assigned, an identifier **MUST NOT** be reused.

If an entry changes meaning functionally:

- Create a new requirement with a new identifier for the updated entry. Do not use the old identifier.
- Preserve the historical identifier for backward compatibility: first remove the old content, then add the field, replacing it with `{your new entry id}`.

Non-functional improvements that do not change intent may retain the existing identifier.

Refer to [PR#443](https://github.com/ossf/security-baseline/pull/443) for additional information and rationale.

## A Single Requirement corresponds to a single Test

Every requirement **MUST** describe exactly one testable behavior. Each requirement should map directly to a single implementation test. If a requirement has multiple obligations that can pass/fail independently, then split them into separate requirements.

An easy way to enforce this is to avoid using “and” in the requirement text.

> **Note:** “Testable” here refers to a requirement being written in a way such that any evaluator, be it humans or machine-automated, can get a definite pass or fail after a check, without any ambiguity while making the judgement. For instance, consider the following requirement: “The project MUST publish a SECURITY.md file”, an evaluator would go to the project repository for this check; either there would be a SECURITY.md file or there won’t.

## Requirements Use Only MUST or MUST NOT

Requirement statements are normative.

Requirements may only contain:

- **MUST**
- **MUST NOT**

Do not use:

- **SHOULD**
- **SHOULD NOT**
- **MAY**
- **Recommended**
- **Preferably**

If guidance is optional or advisory, place it in the Recommendation section instead.

### Example

**Requirement**

> The project **MUST** publish a `SECURITY.md` file.

**Recommendation**

> The `SECURITY.md` file should include contact information and expected response timelines.

## Recommendations Contain Guidance

Recommendations can be used to describe the requirements and should suggest one or more ways to satisfy a particular requirement. These are non-binding suggestions for adopters and evaluators.

**Requirement**

> The project **MUST** prevent the unintentional storage of unencrypted sensitive data in the version control system.

**Recommendation**

> Document how secrets and credentials are managed and used within the project. This should include details on how secrets are stored (e.g., using a secrets management tool), how access is controlled, and how secrets are rotated or updated. Ensure that sensitive information is not hard-coded in the source code or stored in version control systems.

## Avoid using ambiguous qualifiers in requirements

Do not use qualifiers that make a requirement subjective or difficult to evaluate consistently. If a requirement applies universally, state it directly. If it applies only under specific conditions, express those conditions explicitly using **When**, **If**, or **For each**.

Avoid phrases such as:

- While active
- When applicable
- As appropriate
- Where possible
- If feasible
- Normally
- Typically
- In general

The table below shares a few alternate phrases and when to use them:

| Intended meaning | Preferred pattern | Example |
|---|---|---|
| The requirement always applies (Note: When the requirement applies to the entire project) | No qualifier | The project **MUST** publish a LICENSE file. |
| Requirement is applicable under a specific condition | When / If [condition] | When the project publishes a release, the release **MUST** include a changelog. |
| Requirement is applicable to every component (Note: When nothing is exempt from the scope of a requirement with a defined collection) | For each [item] | For each release, the project **MUST** provide a unique version identifier. |

## Use Active Voice in Requirements

Write requirements using active voice.

Active voice is a sentence structure in which the grammatical subject of a sentence performs the action expressed by the verb. It follows the pattern: 

**Subject + Verb + Object**

This may alternatively be described as:

`{CONDITION} {ACTOR} {MUST/NOT} {ACTION}`

**Preferred:**

> The project **MUST** publish a LICENSE file.

**Avoid:**

> A LICENSE file **MUST** be published by the project.

## Prefer Security Outcomes over Implementation Mechanisms

Requirements **SHOULD** describe the security property or outcome that must be achieved rather than prescribing a specific implementation mechanism, unless the mechanism itself is the normative requirement.

**Preferred:**

> The project **MUST** use a delivery mechanism that prevents adversary-in-the-middle attacks.

**Avoid:**

> The project **MUST** use HTTPS to distribute release artifacts.

**Context:** Feedback from the community highlighted cases where implementation-specific words were interpreted more broadly than they were intended (refer [issue#100](https://github.com/ossf/security-baseline/issues/100)).

## Clearly Separate Requirements from Implementation Examples

Requirements define the mandatory behavior that **MUST** be satisfied.

Recommendations provide non-normative examples of one or more ways to satisfy the requirement. Recommendations **MUST NOT** imply that a specific implementation is required unless explicitly stated in the Requirement.

**Requirement**

> The version control system **MUST** require code contributors to assert that they are legally authorized to make their contributions.

**Recommendation**

> Projects may satisfy this requirement using a Developer Certificate of Origin (DCO), Contributor License Agreement (CLA), platform-provided contributor terms, or another equivalent mechanism.

**Context:** The discussion (refer [issue#255](https://github.com/ossf/security-baseline/issues/255)) highlights an instance where recommendations were interpreted as mandatory implementation requirements by the community.
"""
