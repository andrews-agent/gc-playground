## Finding

Warehouse must return the record identified by its requested ID, but existing tests allowed incorrect returned IDs.

## Evidence

In an isolated copy, forcing Warehouse to return ID zero passed the original full suite. The new test rejected that mutation for all three warehouses.

## Selected Candidate

Single-warehouse result coverage: concrete mutation evidence and consistency with existing complete-record batch lookup tests.

## Rejected Candidates

Stock-report batching and arithmetic coverage overlap excluded PRs. Documentation and navigation were current. Loop rewrites were cosmetic; unused-helper removal had weaker justification.

## Change

Added 18 lines to inventory/store_test.go. Changes remain uncommitted; no report file created.

## Validation

make check passed all tests and go vet. gofmt and git diff --check were clean. Make emitted a nonfatal macOS cache-permission diagnostic.

## Risk

Low; tests only.

## Rollback

Remove TestWarehouse from inventory/store_test.go.

## Follow-up

None
