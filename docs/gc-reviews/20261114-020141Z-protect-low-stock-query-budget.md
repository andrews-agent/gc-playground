## Finding

LowStock's output tests did not enforce its constant query budget.

## Evidence

Replacing the batch lookup temporarily with per-item lookups left existing LowStock tests passing; the new test failed with 4 and 6 queries.

## Selected Candidate

Protect LowStock batching: directly measurable, narrowly scoped test coverage.

## Rejected Candidates

StockReport batching and arithmetic coverage overlap open PRs. Removing the unused median helper offered weaker value.

## Change

Added 17 lines in inventory/inventory_test.go covering no matches, one match, partial matches, and all matches, with at most two queries.

## Validation

Ran go mod download: no dependencies to download. Targeted test passed; mutation proof failed as expected. Restored production code, then make check passed all tests and go vet using writable temporary/cache directories. git diff --check passed. No checks remain blocked.

## Risk

Low; test-only change using the existing query counter.

## Rollback

Remove TestLowStockQueryBudget.

## Follow-up

None
