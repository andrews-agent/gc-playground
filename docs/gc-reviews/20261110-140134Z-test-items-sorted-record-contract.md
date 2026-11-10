## Finding

Store.Items promises every item sorted by SKU, but direct store tests did not verify that contract for multiple items.

## Evidence

inventory/store.go documents and implements ascending SKU ordering. The new test passes normally and fails both multi-item cases when the comparator is temporarily reversed.

## Selected Candidate

Close the direct Items contract coverage gap using the existing table-driven test conventions.

## Rejected Candidates

Stock-report batching and arithmetic coverage overlap excluded PRs. Loop rewrites are cosmetic; unused median removal lacks comparable evidence. Documentation and navigation screening found no concrete remaining defect.

## Change

Added 31 lines in inventory/store_test.go covering all identified Items result shapes and comparing complete records. Changes remain uncommitted.

## Validation

make check passed with writable Go caches: all tests and go vet succeeded. git diff --check and gofmt checks passed. Comparator mutation failed as expected and was restored. macOS emitted nonfatal cache warnings.

## Risk

Low: test-only change; production behavior is unchanged.

## Rollback

Revert the added TestItems function and slices import.

## Follow-up

None
