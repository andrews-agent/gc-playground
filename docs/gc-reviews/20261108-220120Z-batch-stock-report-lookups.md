## Finding

StockReport violated the existing batch-query pattern by issuing N+1 queries for N items.

## Evidence

The sample contains five items sharing three warehouses. The regression test failed before the fix with six queries instead of two. Repository-wide search found no other per-item warehouse query loop.

## Selected Candidate

StockReport query batching: directly measurable, locally established, and behavior-preserving.

## Rejected Candidates

Error wrapping and Sum/Mean coverage overlap open PRs. The stale documented threshold is a separate cleanup. Removing the unused median helper offered less concrete value.

## Change

Updated inventory/report.go and inventory/inventory_test.go: 47 insertions, 7 deletions. Nonempty reports use two queries regardless of item count; empty reports retain one. Tests cover exact sorted output, shared warehouses, empty inventory, and missing warehouses without partial output.

## Validation

The regression test failed on the original code. After the fix, make check passed all tests and go vet; git diff --check passed. Writable temporary/cache paths were supplied for the sandbox. make emitted a nonfatal xcrun cache warning.

## Risk

Low. Adds an O(N) ID slice and a warehouse map while preserving item order and error text.

## Rollback

Revert the changes in inventory/report.go and inventory/inventory_test.go.

## Follow-up

None
