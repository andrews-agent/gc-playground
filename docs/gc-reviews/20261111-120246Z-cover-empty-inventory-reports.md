## Finding

Report tests used populated stores and could miss erroneous failures on empty inventories.

## Evidence

Both functions list matching items without requiring a nonempty store. In an isolated copy, returning an error for empty inventories passed the previous inventory suite; all four new cases rejected that mutation.

## Selected Candidate

Close the empty-store coverage gap across both report functions using the existing table-driven test convention.

## Rejected Candidates

Stock-report batching and arithmetic coverage overlap excluded PRs. Documentation links and commands are current. Sum loop changes are cosmetic; unused median removal offered weaker evidence.

## Change

Added one 22-line table-driven test in inventory/inventory_test.go. Product code is unchanged.

## Validation

Ran go mod download: no dependencies to download. make check passed all tests and go vet; make emitted nonfatal macOS cache-write warnings. Mutation verification demonstrated the coverage gap. gofmt and git diff --check passed. No repository AGENTS.md files were present.

## Risk

Low; tests only.

## Rollback

Remove TestReportsEmptyStore.

## Follow-up

None
