## Finding

The existing test allowed lookup buffer reuse to silently overwrite previously returned warehouse maps.

## Evidence

Warehouses currently allocates independent maps. A temporary reused-map mutation passed the original test but failed the strengthened assertion.

## Selected Candidate

Close the warehouse snapshot coverage gap: a concrete invariant already tested for Items.

## Rejected Candidates

Stock report batching and arithmetic tests duplicate excluded open PRs. Removing the unused median helper lacked the required supporting evidence. Reviewed documentation links and command usage appeared consistent.

## Change

Added 14 lines to inventory/store_test.go checking snapshot contents after warehouse replacement and a subsequent subset lookup. Production code remains unchanged.

## Validation

Ran go mod download: no dependencies to download. Targeted test passed; mutation verification distinguished old and new tests. make check passed all tests and go vet with writable Go cache/temp paths. make emitted a nonfatal xcrun cache warning. git diff --check passed.

## Risk

Low; test-only change.

## Rollback

Remove the added snapshot assertion block.

## Follow-up

None
