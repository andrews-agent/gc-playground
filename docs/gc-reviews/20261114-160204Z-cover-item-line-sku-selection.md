## Finding

TestItemLine checked only GEAR-S, leaving the invariant that output describes the requested item unprotected.

## Evidence

Temporarily hardcoding ItemLine’s lookup to gear-s passed the original test but failed four expanded cases. The mutation was restored.

## Selected Candidate

Strengthen successful ItemLine coverage: a concrete, mutation-proven test gap with a small, isolated fix.

## Rejected Candidates

StockReport batching and arithmetic coverage overlap earlier open PRs. Documentation and navigation screening found no further concrete defect.

## Change

Converted TestItemLine into five named table-driven cases in inventory/inventory_test.go; 19 insertions and 6 deletions.

## Validation

Ran go mod download: no external dependencies. Initial make check encountered sandbox cache restrictions; rerunning with writable Go caches passed go test ./... and go vet ./... (exit 0). git diff --check passed. No repository AGENTS.md files were present.

## Risk

Low: tests only; production behavior is unchanged.

## Rollback

Restore the previous TestItemLine implementation.

## Follow-up

None
