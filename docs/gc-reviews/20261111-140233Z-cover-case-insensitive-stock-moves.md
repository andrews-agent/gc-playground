## Finding

Successful Move tests used only uppercase SKUs, leaving canonical SKU preservation unprotected.

## Evidence

Item normalizes lookup casing, and Move uses the canonical returned SKU. A temporary mutation preserving input casing passed all existing inventory tests but failed the new lowercase and mixed-case cases.

## Selected Candidate

Close the case-insensitive Move coverage gap using the existing normalization contract and test conventions.

## Rejected Candidates

Stock-report batching and arithmetic coverage overlap open PRs. Removing the unused median helper offered weaker evidence of practical value. CLI and documentation screening found no clear additional defect.

## Change

Added 22 lines in inventory/inventory_test.go covering three SKU casing variants and checking complete resulting records.

## Validation

Ran go mod download: no dependencies to download. Targeted test passed. Mutation check confirmed the new test detects previously missed corruption; restored production code afterward. make check passed all tests and go vet with writable TMPDIR and GOCACHE. git diff --check passed. macOS cache warnings did not prevent checks from completing.

## Risk

Low; test-only change with no product behavior changes.

## Rollback

Remove TestMoveCaseInsensitiveSKU.

## Follow-up

None
