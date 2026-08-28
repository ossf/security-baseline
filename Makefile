# SPDX-FileCopyrightText: Copyright 2026 The OSPS Baseline Authors
# SPDX-License-Identifier: Apache-2.0

.PHONY: dev
# Compile the devel pages the same way web-publish.yml does, then serve the
# site locally at http://127.0.0.1:4000
dev:
	cd cmd && go run . compile --output ../docs/versions/devel.md \
		--checklist-output ../docs/versions/devel-checklist.md \
		--crosswalk-output ../docs/versions/devel-crosswalk.md
	cd docs && bundle check >/dev/null || (cd docs && bundle install)
	cd docs && bundle exec jekyll serve
