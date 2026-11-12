## Finding

Move trusted destination SKU identity without checking its WarehouseID, silently relocating existing stock.

## Evidence

With BOLT-10@2 stored in warehouse 3, moving BOLT-10 to warehouse 2 overwrote that record's warehouse. The new test failed on original code for quantities 0 and 20.

## Selected Candidate

Destination warehouse conflicts: a demonstrated stock-integrity bug with a localized guard and deterministic regression test.

## Rejected Candidates

Stock-report batching and arithmetic tests overlap excluded open PRs. Removing unused median lacked the required correctness evidence. Documentation and navigation screening found no concrete defect.

## Change

Validate the existing destination's warehouse before mutation, reuse its loaded quantity, and document rejection. Repository-wide search found one affected transfer path. Changed two files: 26 insertions and 2 deletions.

## Validation

No repository AGENTS.md files exist. Ran go mod download successfully; no external dependencies. Regression test failed before the fix; all move tests passed afterward. Initial make check encountered sandbox cache restrictions. Rerunning with writable GOCACHE, GOTMPDIR, and TMPDIR passed go test ./... and go vet ./...; make emitted a nonfatal xcrun cache warning. git diff --check passed.

## Risk

Previously accepted conflicting transfers now return an error; valid transfers retain existing behavior.

## Rollback

Revert the changes to inventory/store.go and inventory/inventory_test.go.

## Follow-up

None
