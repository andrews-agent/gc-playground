## Finding

The error-identity tests omitted LowStock, allowing its error chain to break undetected.

## Evidence

Changing LowStock's wrapping from %w to %v passed the previous full suite. The added case fails on that mutation.

## Selected Candidate

Complete the existing error-identity coverage: an established caller contract with a deterministic mutation check.

## Rejected Candidates

Stock-query batching and arithmetic coverage overlap excluded PRs. Removing the unused median helper offers less concrete value; movement changes require broader behavioral decisions.

## Change

Added one five-line case in inventory/inventory_test.go. Production code is unchanged; changes remain uncommitted.

## Validation

Mutation check confirmed old-suite pass and new-case failure. Restored production code, then make check passed tests and go vet. git diff --check passed. Used a writable Go cache; make emitted a nonfatal macOS cache warning.

## Risk

Low: test-only change using the existing fixture and assertion.

## Rollback

Remove the added test-table case.

## Follow-up

None
