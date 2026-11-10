## Finding

Store.Move accepted negative quantities, increasing source stock and reducing or creating negative destination stock.

## Evidence

Both regression cases failed before the fix: source stock increased from 120 to 121; a new destination received -1. Both pass after the fix.

## Selected Candidate

Missing lower-bound validation in Store.Move: a demonstrated stock mutation bug with one owning entry point.

## Rejected Candidates

StockReport batching and arithmetic test improvements overlap excluded open PRs. Documentation and navigation screening found no concrete remaining issue. Removing unused median offered less value.

## Change

Six added lines across two files: a negative-quantity guard, its documentation, and two regression cases. Repository search found no other transfer implementation.

## Validation

Targeted regression and zero-quantity tests passed. make check passed with writable GOCACHE and TMPDIR overrides, running all tests and go vet. git diff --check passed. macOS emitted a nonfatal cache warning.

## Risk

Low. Negative moves now return an error; existing zero and positive move behavior remains covered.

## Rollback

Revert the six added lines in inventory/store.go and inventory/inventory_test.go.

## Follow-up

None
