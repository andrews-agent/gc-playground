## Finding

No tests enforced the documented Queries counter contract.

## Evidence

Removing each of the four counter increments independently passed the existing suite but failed the new test.

## Selected Candidate

Missing query-counter coverage: an explicit local contract with deterministic mutation evidence.

## Rejected Candidates

Report batching and arithmetic coverage overlap excluded PRs. SKU normalization lacks a clear contract. Loop styling and unused-helper removal offered weaker evidence.

## Change

Added inventory/store_test.go with 16 cases in 40 lines. Changes remain uncommitted; no report file created.

## Validation

make check passed all tests and go vet. All four mutation checks distinguished existing from new coverage. gofmt and whitespace checks passed. Make emitted a nonfatal sandbox cache warning.

## Risk

Low; tests only.

## Rollback

Remove inventory/store_test.go.

## Follow-up

None
