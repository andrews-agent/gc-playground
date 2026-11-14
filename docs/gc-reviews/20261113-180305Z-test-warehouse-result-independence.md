## Finding

Warehouse lookup tests checked contents but did not protect result independence.

## Evidence

Warehouses constructs a separate map. A temporary mutation returning internal storage passed all existing inventory tests but failed the new test; the mutation was restored.

## Selected Candidate

Close the demonstrated warehouse map isolation coverage gap with one focused test.

## Rejected Candidates

Stock-report batching and arithmetic tests overlap excluded PRs. Documentation navigation and commands were already correct. Other behavior changes lacked an established contract.

## Change

Added one 30-line test in inventory/store_test.go covering returned-map replacement, deletion, and independence from store additions.

## Validation

Ran go mod download: no external dependencies. make check passed all tests and go vet with a writable GOCACHE after the default cache was blocked by sandbox permissions. git diff --check passed. Mutation testing confirmed the new assertion detects a previously undetected failure.

## Risk

Low; test-only change.

## Rollback

Remove TestWarehousesResultIsIndependent.

## Follow-up

None
