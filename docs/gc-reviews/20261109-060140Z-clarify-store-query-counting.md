## Finding

The Store documentation contradicted its query counter implementation.

## Evidence

A runtime probe confirmed AddWarehouse counts 0 queries, AddItem counts 1, Move counts 2, and a failed Warehouse lookup counts 1. All counter increments occur in the four documented lookup methods.

## Selected Candidate

Correct query-counter documentation: directly verifiable, complete, and behavior-preserving.

## Rejected Candidates

Stock-report batching and arithmetic coverage overlap excluded PRs. Loop rewrites are cosmetic; changes to Move validation lack a sufficiently explicit contract.

## Change

Updated one comment in inventory/store.go: 3 insertions and 2 deletions. Searched all query-count references; no other current claims required correction. Changes remain uncommitted.

## Validation

make check passed all tests and go vet, with a nonfatal macOS cache warning. Runtime probe confirmed counter semantics. git diff --check passed; reviewed the complete diff.

## Risk

Minimal; documentation only.

## Rollback

Restore the previous Store comment.

## Follow-up

None
