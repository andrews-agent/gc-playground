## Finding

The warehouses callback captured the parent testing.T, violating the test harness invariant that subtest assertions report against the active subtest.

## Evidence

A temporary forced assertion failure made the original warehouses subtest report PASS while its parent failed. With the fix, warehouses correctly reported FAIL. The temporary mutation was restored.

## Selected Candidate

Fix assertion ownership in TestMissingLookupsInEmptyStore: a concrete, reproducible tooling defect with a five-line replacement.

## Rejected Candidates

Stock-report batching and arithmetic coverage overlap excluded PRs. Documentation and navigation were current. Unused median removal lacked sufficient supporting evidence; loop styling was cosmetic.

## Change

Updated the lookup callback signature and invocations in inventory/store_test.go to receive the active testing.T. Screening all repository tests found no other instance.

## Validation

Ran go mod download: no external dependencies. Initial make check was blocked by sandbox cache permissions. With writable GOCACHE and TMPDIR, make check passed all tests and go vet. git diff --check passed. Before/after failure-attribution experiment confirmed the fix.

## Risk

Low: test harness only.

## Rollback

Revert the callback signature and argument changes in inventory/store_test.go.

## Follow-up

None
