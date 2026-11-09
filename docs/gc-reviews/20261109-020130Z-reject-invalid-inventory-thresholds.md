## Finding

The inventory low command discarded strconv.Atoi errors, violating the existing CLI convention that invalid numeric arguments fail without report output.

## Evidence

Before the fix, regression tests showed invalid text returning the zero-threshold report and overflow returning the full report, both with exit code 0.

## Selected Candidate

Fix discarded threshold parsing errors: a directly reproduced bug with a small, local correction. Repository-wide search found one instance.

## Rejected Candidates

Low-stock documentation and stock-report batching overlap excluded PRs. The log-only stock-report test overlaps the batching work. Removing unused median lacked a demonstrated correctness benefit.

## Change

Check the parsing error, print a diagnostic, and exit 2. Add five CLI cases following the existing subprocess-test convention. Changes remain uncommitted.

## Validation

Regression tests failed before the fix and passed afterward. make check completed successfully: all package tests and go vet passed. git diff --check passed. macOS emitted a nonfatal cache-permission diagnostic.

## Risk

Invalid thresholds now fail explicitly; default, zero, and negative thresholds retain their tested behavior.

## Rollback

Revert cmd/inventory/main.go and remove cmd/inventory/main_test.go.

## Follow-up

None
