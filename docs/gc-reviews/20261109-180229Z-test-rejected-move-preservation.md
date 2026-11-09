## Finding

Move failure paths lacked coverage proving that rejected transfers leave inventory unchanged.

## Evidence

Store.Move validates missing items, missing warehouses, and insufficient stock before writing. Removing the stock guard passed the original inventory suite; the new test caught the accepted transfer and corrupted quantities.

## Selected Candidate

Strengthen rejected-move tests: an established invariant with deterministic mutation evidence and no product changes.

## Rejected Candidates

Stock batching and arithmetic coverage overlap excluded PRs. Negative-quantity validation lacks an explicit contract. Documentation navigation was intact; loop rewrites were cosmetic.

## Change

Added 29 lines in inventory/inventory_test.go covering all three rejection paths. Production code is unchanged; changes remain uncommitted.

## Validation

make check passed all tests and go vet with writable Go cache paths. git diff --check passed. Mutation was restored after demonstrating test failure. Make emitted a nonfatal sandbox cache warning.

## Risk

Low: test-only change.

## Rollback

Remove TestRejectedMovePreservesItems and the slices import.

## Follow-up

None
