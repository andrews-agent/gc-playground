## Finding

Inventory wrappers discarded error identity, defeating the CLI's errors.Is check and documented unknown-item response.

## Evidence

All six new regression cases failed before the fix and passed afterward. LowStock already preserves error identity with %w.

## Selected Candidate

Error identity preservation: an established caller contract, six concrete violations, and deterministic regression tests.

## Rejected Candidates

Report query batching and the stale documented threshold are separate cleanup classes. Arithmetic coverage overlaps the excluded PR; loop styling is cosmetic.

## Change

Updated all six affected wrappers in inventory/report.go and inventory/store.go; added table-driven tests. Changes remain uncommitted.

## Validation

Regression tests failed on the original code. make check passed go test ./... and go vet ./..., with a nonfatal macOS cache warning. gofmt and git diff --check passed. No fmt.Errorf wrappers using %v remain.

## Risk

Low. Error text remains unchanged; errors.Is now recognizes wrapped causes as intended.

## Rollback

Revert the changes in the three modified inventory files.

## Follow-up

None
