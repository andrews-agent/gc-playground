## Finding

Batch lookup tests checked error identity but ignored results, allowing partial warehouse maps on failure to pass.

## Evidence

Store.Warehouses explicitly returns nil when any ID is missing. A temporary mutation returning accumulated results passed the original tests and failed all five strengthened cases.

## Selected Candidate

Close the batch lookup error-result coverage gap with eight added lines across existing tests.

## Rejected Candidates

Stock-report batching and arithmetic tests overlap the listed open PRs. Removing unused median code offered less concrete value; other behavior changes lacked an established contract.

## Change

Added nil-result assertions to missing-ID and empty-store batch lookup tests in inventory/store_test.go.

## Validation

Ran go mod download: no dependencies to download. Targeted tests passed. Mutation verification distinguished original and strengthened tests; production code was restored. make check passed all tests and go vet with a writable GOCACHE, despite a nonfatal macOS cache warning. git diff --check passed.

## Risk

Low; test-only change preserves the implementation's existing error-result behavior.

## Rollback

Revert the assertions in inventory/store_test.go.

## Follow-up

None
