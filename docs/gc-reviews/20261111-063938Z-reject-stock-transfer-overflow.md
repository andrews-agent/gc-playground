## Finding

Store.Move could violate stock conservation by overflowing destination quantity into a negative integer.

## Evidence

Before the fix, two regression cases accepted overflowing transfers and produced destination quantity -9223372036854775808. Both also changed source stock. Repository-wide search found one transfer implementation.

## Selected Candidate

Destination overflow: a reproducible correctness bug with a local fix and decisive regression tests.

## Rejected Candidates

Stock-report batching and arithmetic coverage overlap excluded PRs. Unused median removal lacks comparable evidence. Documentation and navigation screening found no concrete remaining defect.

## Change

Added a capacity guard before either inventory write, updated the method comment, and added four boundary cases across two files.

## Validation

No repository AGENTS.md files exist. Ran go mod download before checks; no external dependencies were needed. The targeted regression test failed before the fix and passed afterward. GOCACHE="$PWD/../tmp/go-cache" make check passed all tests and go vet despite nonfatal macOS cache warnings. git diff --check passed.

## Risk

Low. Transfers exceeding the platform's maximum int now fail explicitly; boundary-valid transfers remain supported.

## Rollback

Revert the changes to inventory/store.go and inventory/inventory_test.go.

## Follow-up

None
