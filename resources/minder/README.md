# OSPS Minder Profiles

[Minder](https://github.com/mindersec/minder) is an Open Source Supply Chain
Security plaform part of the OpenSSF ecosystem. A Minder profile is a security
policy that groups together rules that are constantly evaluating against entities
in your software project.

This directory contains three [Minder](https://github.com/mindersec/minder) 
profiles that any project can use to monitor its compliance with the 
Open Source Project Security Baseline.

Each file is designed to be applied incrementally: Projects at Level 1 
(sandbox or similar) should only apply the Level 1 file, projects on Level 2 
(incubating or similar) should apply Level 1 and Level 2 and at Level 3
you should apply all files.

## Installing

To install the profiles and monitor a project using the public Minder instance,
follow these steps:

1. Install the Minder CLI
1. Login using your GitHub account
1. Enroll the GitHub provider in your project
1. Apply the profile and all the OSPS rule types.

### Detailed instructions

These are installation steps in full detail. If you run into a problem, see
the section on getting help below. The following commands need to be run
in a computer with a web browser installed.

#### 1. Install the Minder CLI

Follow the 
[instructions on the Minder README](https://github.com/mindersec/minder#installation) 
to install the `minder` CLI in your computer.

#### 2. Log In With Your GitHub Account

Run the following commands to authenticate:

```
minder auth login
```

Follow the authentication flow to log in. This will create your first (empty)
project.

#### 3. Enroll the GitHub Provider in your Project.

```bash
minder provider enroll
```

Follow the GitHub authentication flow to connect minder to your GitHub organization

#### 4. Apply the OSPS Profile(s) and Rule Types

Finally, apply the appropiate profiles according to you project's maturity level
and the OSPS rule types. 

Clone the OSPS Baseline repository and apply the minder resource files: 

```bash
git clone git@github.com:ossf/security-baseline.git

cd security-baseline

minder ruletype apply -f resources/minder/rules/

minder profile apply -f resources/minder/osps-baseline-level1.yaml
```

Remember to install the profile for your level and all the previous ones.
For example, if your project is on maturity level 3 (graduated or similar), 
apply the Level 3 profile, but also 1 & 2.

## Help and Support

If you have questions or need help, please talk to the Minder maintainers in
[#minder in OpenSSF Slack](https://openssf.slack.com/archives/C07SP9RSM2L) 
([get an invite here](https://openssf.org/getinvolved/) if you are not on it
yet). We would love to hear your experience as the project evolves to help you
monitor and remediate the OSPS Baseline checks.
