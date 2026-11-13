## Finding

Existing report error tests discarded output, allowing partial output regressions to pass.

## Evidence

All four error paths explicitly return empty strings. Mutating them to return nonempty output left existing error tests passing but failed all four new cases.

## Selected Candidate

Incomplete report error assertions: a concrete, consistently implemented invariant with demonstrably weak coverage.

## Rejected Candidates

Stock-report batching and arithmetic tests overlap excluded PRs. Loop styling is cosmetic; removing the unused median helper offered less practical value.

## Change

Added 25 lines in inventory/inventory_test.go covering ItemLine, StockReport, and LowStock failure output. Production behavior is unchanged.

## Validation

go mod download completed: no external dependencies. make check passed tests and go vet with writable cache paths; macOS emitted a nonfatal cache warning. Mutation testing distinguished the new assertions from existing coverage. gofmt and git diff --check passed. No AGENTS.md files exist.

## Risk

Low; test-only change.

## Rollback

Remove TestReportsDiscardOutputOnError.

## Follow-up

None
