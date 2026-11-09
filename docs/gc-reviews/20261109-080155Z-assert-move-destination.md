## Finding

TestMoveInBackground checked only source quantity, allowing broken destination persistence to pass.

## Evidence

Store.Move must save the transferred quantity in the destination warehouse. Three isolated mutations—missing destination, wrong warehouse, and wrong quantity—passed the original suite and failed the strengthened test.

## Selected Candidate

Incomplete move-result assertions: concrete coverage gap, deterministic evidence, and a small test-only fix. Repository-wide inspection found one successful-move test.

## Rejected Candidates

Query batching, query-count documentation, and arithmetic coverage overlap excluded PRs. Loop styling is cosmetic; SKU normalization lacks a clear contract.

## Change

Updated inventory/inventory_test.go with complete source and destination assertions: 11 insertions, 2 deletions. Changes remain uncommitted.

## Validation

Targeted test, make check, go test -race ./..., and git diff --check passed. All three mutation checks distinguished the original and updated suites. Make emitted a nonfatal sandbox cache warning.

## Risk

Low; tests only.

## Rollback

Restore the previous assertions in TestMoveInBackground.

## Follow-up

None
