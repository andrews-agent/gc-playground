## Finding

Max promises to panic on empty input, but existing tests covered only nonempty slices.

## Evidence

A temporary mutation returning zero for empty input passed existing Max tests and failed both new cases. The mutation was reverted.

## Selected Candidate

Missing Max panic-contract coverage: explicit, locally documented, and demonstrably distinguishable.

## Rejected Candidates

Stock-report queries, move assertions, Store documentation, and Sum/Mean coverage overlap listed open PRs. Negative move validation lacks an explicit established contract. Navigation links resolve.

## Change

Added 20 lines in stats/stats_test.go covering both empty-input forms. Changes remain uncommitted.

## Validation

make check passed with writable TMPDIR and GOCACHE after default cache permissions blocked the first attempt. All package tests and go vet passed. git diff --check and gofmt checks passed.

## Risk

Low; tests only.

## Rollback

Remove TestMaxEmptyPanics from stats/stats_test.go.

## Follow-up

None
