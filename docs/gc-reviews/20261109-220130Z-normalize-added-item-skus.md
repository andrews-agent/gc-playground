## Finding

Successfully added lowercase SKUs could not be retrieved because storage and lookup used different keys.

## Evidence

Item uppercases lookup keys; AddItem previously stored them unchanged. The new test failed for lowercase and mixed-case additions before the fix and passed afterward.

## Selected Candidate

SKU normalization: a reproduced correctness bug with a one-line fix at the storage boundary. Other item writes derive their keys from stored SKUs.

## Rejected Candidates

Stock-report batching and arithmetic test improvements overlap excluded PRs. Dead-code removal lacked comparable correctness evidence; loop-style changes were cosmetic.

## Change

Added uppercase normalization, a method comment, and table-driven regression coverage. Two files changed, 28 lines added; changes remain uncommitted.

## Validation

Targeted regression passed. make check passed with a writable GOCACHE, running all tests and go vet. git diff --check and gofmt checks passed. Apple's make launcher emitted a nonfatal cache warning.

## Risk

Low. Added SKUs now use uppercase consistently; case variants resolve to the same storage key.

## Rollback

Revert the changes in inventory/store.go and inventory/inventory_test.go.

## Follow-up

None
