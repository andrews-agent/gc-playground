## Finding

Batch lookup tests checked query counts without verifying returned warehouse records.

## Evidence

Temporarily zeroing returned warehouse IDs passed the existing suite. The new tests failed for every nonempty case. The mutation was then reverted.

## Selected Candidate

Strengthen Warehouses result coverage: a documented contract with a demonstrated test gap.

## Rejected Candidates

Stock-report batching and arithmetic coverage overlap excluded PRs. Negative-move validation lacks an explicit contract. Documentation links and commands were current; loop rewrites were cosmetic.

## Change

Changed only inventory/store_test.go: 32 insertions and 1 deletion. Product behavior is unchanged; changes remain uncommitted.

## Validation

make check passed with writable TMPDIR and GOCACHE: all tests and go vet succeeded. git diff --check passed. The initial default-cache run was blocked by sandbox permissions.

## Risk

Low; test-only change using existing warehouse records and standard-library map comparison.

## Rollback

Revert the changes to inventory/store_test.go.

## Follow-up

None
