## Finding

AddItem rejection tests checked error identity without checking that inventory remained unchanged.

## Evidence

AddItem validates the warehouse before writing. A temporary write-on-error mutation passed the existing suite but failed all three new cases; the mutation was removed.

## Selected Candidate

Strengthen rejected-add state preservation coverage, following the existing rejected-move test convention.

## Rejected Candidates

Stock-report batching and arithmetic tests overlap excluded PRs. Documentation and navigation were current. Loop rewrites lacked correctness value.

## Change

Added 16 lines in inventory/inventory_test.go. Product code is unchanged; changes remain uncommitted.

## Validation

make check passed all tests and go vet with writable temporary/cache paths, despite a macOS cache warning. gofmt and git diff --check were clean.

## Risk

Low: test-only change using existing fixtures and assertions.

## Rollback

Remove TestRejectedAddItemPreservesItems.

## Follow-up

None
