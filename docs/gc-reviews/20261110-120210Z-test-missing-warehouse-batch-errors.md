## Finding

The missing-ID batch lookup test checked query counts but ignored errors, leaving part of the documented ErrNotFound contract unprotected.

## Evidence

In an isolated copy, returning success when the first requested warehouse was missing passed the original full suite. The new tests rejected that mutation.

## Selected Candidate

Batch lookup failure coverage: an established contract with a demonstrated test gap.

## Rejected Candidates

Stock-report batching and arithmetic coverage overlap excluded PRs. Documentation and navigation were current. Loop rewrites and unused-helper removal lacked comparable evidence.

## Change

Added 20 lines in inventory/store_test.go covering all missing-ID positions. Changes remain uncommitted.

## Validation

make check passed all tests and go vet. Mutation verification, gofmt, and git diff --check passed. Make emitted a nonfatal sandbox cache diagnostic.

## Risk

Low; tests only.

## Rollback

Remove TestWarehousesMissingID and its errors import.

## Follow-up

None
