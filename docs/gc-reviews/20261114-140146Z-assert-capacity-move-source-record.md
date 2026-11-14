## Finding

TestMoveDestinationCapacity checked only source quantity, leaving the invariant that moves preserve source SKU, name, and warehouse unprotected.

## Evidence

Other successful-move tests compare complete records. A temporary mutation corrupting the source name passed the original capacity test and failed the strengthened assertion.

## Selected Candidate

Close the demonstrated partial-record assertion gap with a focused test-only change.

## Rejected Candidates

Stock-report batching and arithmetic tests overlap the listed open PRs. Documentation and navigation checks revealed no current defect; stylistic changes lacked sufficient justification.

## Change

Updated inventory/inventory_test.go: six insertions and two deletions. Searched all successful-move assertions; the other cases already compare complete source records.

## Validation

Ran go mod download before checks; no dependencies required downloading. Mutation verification distinguished old and new assertions, then production code was restored. make check passed all tests and go vet using a writable GOCACHE. git diff --check passed. A nonfatal macOS cache warning did not prevent checks.

## Risk

Low; test-only change with no product behavior changes.

## Rollback

Revert the assertion changes in inventory/inventory_test.go.

## Follow-up

None
