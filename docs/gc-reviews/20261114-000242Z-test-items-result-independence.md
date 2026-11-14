## Finding

Items tests checked values and sorting but missed result independence.

## Evidence

A temporary implementation reusing a result buffer passed the existing TestItems but failed the new test. The mutation was restored.

## Selected Candidate

Close the Items result-independence coverage gap, following an established repository pattern.

## Rejected Candidates

Stock-report batching and arithmetic coverage overlap excluded PRs. Removing unused median lacked sufficient value.

## Change

Added one 22-line test in inventory/store_test.go. Product code is unchanged.

## Validation

go mod download reported no dependencies. Targeted tests, make check (all tests and go vet), and git diff --check passed. Go checks used a writable temporary cache after default-cache access failed. make emitted a nonfatal xcrun cache warning.

## Risk

Low; test-only change.

## Rollback

Remove TestItemsResultIsIndependent.

## Follow-up

None
